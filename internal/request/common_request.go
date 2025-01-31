package request

import (
	"fmt"
	"news/internal/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	LocEn      = "en"
	LocUk      = "uk"
	DefaultLoc = LocEn
)

// LocInWhiteList Перевіряє що передана користувачем локаль дозволена.
func LocInWhiteList(locale string) bool {
	if locale == LocEn {
		return true
	}
	if locale == LocUk {
		return true
	}
	return false
}

// TokenFromRequest Повертає jwt токен з заголовків запиту.
func TokenFromRequest(c *fiber.Ctx) string {
	reqToken := c.GetReqHeaders()["Authorization"][0]
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	return reqToken
}

// ClaimsFromToken Повертає claims з jwt-токена.
func ClaimsFromToken(jwtString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.NewEnv().JwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, nil
}

// CurrentUserId Повертає ID користувача з jwt-токена.
func CurrentUserId(c *fiber.Ctx) (int, error) {
	reqToken := TokenFromRequest(c)
	claims, err := ClaimsFromToken(reqToken)

	if err != nil {
		return 0, err
	}

	userIdRaw := claims["sub"]
	return int(userIdRaw.(float64)), nil
}
