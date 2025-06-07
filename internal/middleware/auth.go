package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sarff/go-robotdreams-diplom/internal/utils"
	log "github.com/sarff/iSlogger"
)

func AuthRequired(secret string) fiber.Handler {
	// TODO: need implement Get UserID from Token
	return func(c fiber.Ctx) error {
		auth := c.Get("X-User-Token")
		log.Debug("auth: %v", auth)
		if auth == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing auth header",
			})
		}
		//authParts := strings.Split(auth, " ")
		//log.Debug("authParts: %v", authParts)
		//if len(authParts) != 2 {
		//	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		//		"error": "invalid auth header",
		//	})
		//}

		claims, err := utils.ValidateToken(auth, secret)
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
