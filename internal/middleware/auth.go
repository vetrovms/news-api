package middleware

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"news/internal/config"
	"news/internal/controllers"
	myerrors "news/internal/errors"
	"news/internal/request"
	"news/internal/response"
	"strconv"
	"strings"
	"time"

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
	reqToken := request.TokenFromRequest(c)
	reader := strings.NewReader("jwt=" + reqToken)
	request, err := http.NewRequest("POST", config.NewEnv().RetrospectiveUrl, reader)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	var retroResp RetrospectiveResponse
	err = json.Unmarshal(body, &retroResp)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	if !retroResp.Data.Result {
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

		// id, err := c.ParamsInt("id") // неприємний сюрприз: на рівні middleware немає доступу до параметрів шляху
		pathParts := strings.Split(c.Path(), "/")
		id, err := strconv.Atoi(pathParts[4])
		if err != nil {
			r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

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
