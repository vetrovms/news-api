package middleware

import (
	"context"
	"news/internal/config"
	"news/internal/controllers"
	myerrors "news/internal/errors"
	"news/internal/request"
	"news/internal/response"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Protected Посередник авторизації.
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(config.NewEnv().JwtSecretKey)},
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
	})
}

// jwtError Перевірка авторизації.
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == myerrors.WrongJWT {
		r := response.NewResponse(fiber.StatusBadRequest, myerrors.WrongJWT, nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}
	r := response.NewResponse(fiber.StatusUnauthorized, myerrors.ExpiredJWT, nil)
	return c.Status(fiber.StatusUnauthorized).JSON(r)
}

// jwtSuccess Ретроспективна перевірка токена.
func jwtSuccess(c *fiber.Ctx) error {
	client := resty.New()

	reqToken := request.TokenFromRequest(c)
	reader := strings.NewReader(
		"jwt=" + reqToken + "&client_id=" + config.NewEnv().ClientId + "&client_secret=" + config.NewEnv().ClientSecret,
	)

	res, err := client.R().
		SetBody(reader).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(&RetrospectiveResponse{}).
		Post(config.NewEnv().RetrospectiveUrl)

	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	if !res.Result().(*RetrospectiveResponse).Data.Result {
		r := response.NewResponse(fiber.StatusBadRequest, myerrors.InvalidUser, nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	return c.Next()
}

// RetrospectiveResponse Структура відповіді сервера авторизації.
type RetrospectiveResponse struct {
	Code   int
	Errors []string
	Data   RetrospectiveResponseData
}

// RetrospectiveResponseData Структура даних відповіді сервера авторизації.
type RetrospectiveResponseData struct {
	Result bool
}

// Config Конфігурація посередника перевірки авторства статті.
type Config struct {
	Filter  func(c *fiber.Ctx) bool
	Service controllers.NewsService
}

// CheckAuthor Посередник перевірки авторства статті.
func CheckAuthor(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// id := c.Params("id") // на рівні middleware немає доступу до параметрів шляху
		pathParts := strings.Split(c.Path(), "/")
		id := pathParts[4]

		curUserId, err := request.CurrentUserId(c)
		if err != nil {
			r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		article, err := config.Service.OneUnscoped(ctx, c, id)
		if err != nil {
			r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		if article.UserId != curUserId {
			r := response.NewResponse(fiber.StatusForbidden, myerrors.InvalidUser, nil)
			return c.Status(fiber.StatusForbidden).JSON(r)
		}

		return c.Next()
	}
}
