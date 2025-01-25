package middleware

import (
	"news/internal/config"
	myerrors "news/internal/errors"
	"news/internal/response"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Protected Посередник авторизації.
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.NewEnv().JwtSecretKey)},
		ErrorHandler: jwtError,
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
