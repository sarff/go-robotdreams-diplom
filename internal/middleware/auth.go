package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sarff/go-robotdreams-diplom/internal/utils"
	log "github.com/sarff/iSlogger"
)

func AuthRequired(secret string) fiber.Handler {
	return func(c fiber.Ctx) error {
		auth := c.Get("X-User-Token")
		log.Debug("auth:", "X-User-Token", auth)
		if auth == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing auth header",
			})
		}

		claims, err := utils.ValidateToken(auth, secret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}
		log.Debug("claims:", "value", claims)

		userID, ok := claims["user_id"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token payload",
			})
		}
		log.Debug("claims:", "userID", userID)
		c.Locals("userID", userID)
		return c.Next()
	}
}
