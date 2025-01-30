package middleware

import (
	"encoding/json"
	"io"
	"net/http"
	"news/internal/config"
	myerrors "news/internal/errors"
	"news/internal/response"
	"strings"

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
	reqToken := c.GetReqHeaders()["Authorization"][0]
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

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

type RetrospectiveResponse struct {
	Code   int
	Errors []string
	Data   RetrospectiveResponseData
}

type RetrospectiveResponseData struct {
	Result bool
}
