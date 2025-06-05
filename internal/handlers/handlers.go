package handlers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/sarff/go-robotdreams-diplom/internal/services"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func InitFiber(services *services.Services) *fiber.App {
	// Initialize handlers
	authHandler := NewAuthHandler(services.Auth)
	//chatHandler := NewChatHandler(services.Chat)
	//wsHandler := NewWSHandler(hub, chatService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "Internal Server Error"

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				message = e.Message
			}

			return c.Status(code).JSON(fiber.Map{
				"error": message,
			})
		},
	})

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool { return true },
		AllowMethods: []string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodOptions,
		},
		AllowHeaders: []string{"Origin, Content-Type, Accept, Authorization"},
	}))

	app.Get("/docs", func(c fiber.Ctx) error {
		return c.SendFile("./static/index.html")
	})

	//app.Get("/docs/swagger.json", func(c fiber.Ctx) error {
	//	return c.Redirect("/swagger/doc.json", fiber.StatusMovedPermanently)
	//})

	app.Get("/swagger/*", adaptor.HTTPHandler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // URL для swagger.json
	)))

	// Routes
	api := app.Group("/api/v1")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	//auth.Get("/profile", middleware.AuthRequired(cfg.JWTSecret), authHandler.GetProfile)

	// Chat routes
	//chat := api.Group("/chat", middleware.AuthRequired(cfg.JWTSecret))
	//chat.Post("/rooms", chatHandler.CreateRoom)
	//chat.Get("/rooms", chatHandler.GetRooms)
	//chat.Get("/rooms/:roomID/messages", chatHandler.GetMessages)
	//chat.Post("/messages", chatHandler.SendMessage)

	// WebSocket route
	//app.Get("/ws", middleware.WSAuthRequired(cfg.JWTSecret), wsHandler.HandleWebSocket())

	// Health check
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	return app
}
