package server

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
	"github.com/sarff/go-robotdreams-diplom/internal/handlers"
	"github.com/sarff/go-robotdreams-diplom/internal/middleware"
	"github.com/sarff/go-robotdreams-diplom/internal/services"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Server struct {
	app *fiber.App
	cfg *config.Config
}

func NewServer(cfg *config.Config, services *services.Services) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: globalErrorHandler,
	})

	setupCors(app)

	setupRoutes(app, services, cfg)

	return &Server{
		app: app,
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	return s.app.Listen(":" + s.cfg.Listen.Port)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}

func globalErrorHandler(c fiber.Ctx, err error) error {
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
}

func setupCors(app *fiber.App) {
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
}

func setupRoutes(app *fiber.App, services *services.Services, cfg *config.Config) {
	// Static files and docs
	app.Use("/docs", static.New("./static"))

	// Stoplight
	app.Get("/openapi.yaml", func(c fiber.Ctx) error {
		return c.SendFile("./openapi.yaml")
	})

	// Swagger
	app.Get("/swagger/*", adaptor.HTTPHandler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)))

	// Health check
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	api := app.Group("/api/v1")

	authHandler := handlers.NewAuthHandler(services.Auth)
	chatHandler := handlers.NewChatHandler(services.Chat)
	wsHandler := handlers.NewWSHandler(services.WS, services.Chat)

	setupAuthRoutes(api, authHandler, cfg)
	setupChatRoutes(api, chatHandler, cfg)
	setupWebSocketRoutes(api, wsHandler, cfg)
}

func setupAuthRoutes(api fiber.Router, handler *handlers.AuthHandler, cfg *config.Config) {
	auth := api.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Get("/profile", handler.GetProfile, middleware.AuthRequired(cfg.JWT.Secret))
}

func setupChatRoutes(api fiber.Router, handler *handlers.ChatHandler, cfg *config.Config) {
	chat := api.Group("/chat")
	chat.Post("/rooms", handler.CreateRoom, middleware.AuthRequired(cfg.JWT.Secret))
	chat.Get("/rooms", handler.GetRooms, middleware.AuthRequired(cfg.JWT.Secret))
	chat.Get("/id/:roomID", handler.FindRoomByID, middleware.AuthRequired(cfg.JWT.Secret))
	chat.Get("/rooms/:roomName", handler.FindRoomByName, middleware.AuthRequired(cfg.JWT.Secret))
	chat.Get("/rooms/:roomID/messages", handler.GetMessages, middleware.AuthRequired(cfg.JWT.Secret))
	chat.Post("/messages", handler.SendMessage, middleware.AuthRequired(cfg.JWT.Secret))
}

func setupWebSocketRoutes(api fiber.Router, handler *handlers.WebsocketHandler, cfg *config.Config) {
	// TODO:
	api.Get("/ws", handler.HandleWebSocket, middleware.WSAuthRequired())
}
