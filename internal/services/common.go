package services

import (
	"news/internal/request"

	"github.com/gofiber/fiber/v2"
)

// locale Повертає локаль.
func locale(c *fiber.Ctx) string {
	locale := c.Query("locale")
	if !request.LocInWhiteList(locale) {
		locale = request.DefaultLoc
	}
	return locale
}
