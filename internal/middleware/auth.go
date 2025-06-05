package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/sarff/go-robotdreams-diplom/internal/utils"
	log "github.com/sarff/iSlogger"
)

func AuthRequired(secret string) fiber.Handler {
	// TODO: need implement Get UserID from Token
	return func(c fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing auth header",
			})
		}
		authParts := strings.Split(auth, " ")
		if len(authParts) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid auth header",
			})
		}

		claims, err := utils.ValidateToken(authParts[1], secret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}
		log.Debug("claims: %v", claims)
		c.Locals("userID", claims)
		return c.Next()
	}
}
