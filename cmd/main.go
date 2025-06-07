// @title Shat API
// @version 0.001.001
// @description Boom boom — and into production.

// @host 127.0.0.1
// @BasePath /

// @securityDefinitions.apiKey UserTokenAuth
// @in header
// @name X-User-Token

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sarff/go-robotdreams-diplom/internal/clients"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
	"github.com/sarff/go-robotdreams-diplom/internal/server"
	"github.com/sarff/go-robotdreams-diplom/internal/services"
	log "github.com/sarff/iSlogger"

	_ "github.com/sarff/go-robotdreams-diplom/docs"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found, using environment variables")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	logConfig := log.DefaultConfig().
		WithDebug(cfg.IsDebug).
		WithJSONFormat(cfg.LogJsonFormat)

	if err = log.Init(logConfig); err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	defer log.Close()

	ctx := context.Background()

	// Clients (DB,)
	clnts, err := clients.NewClients(ctx, cfg)
	if err != nil {
		log.Error("Failed to create clients: %v", err)
		panic(err)
	}

	// Services (Auth, Chat, WS):
	srvcs, err := services.NewServices(cfg, clnts)

	httpServer := server.NewServer(cfg, srvcs)
	go func() {
		if err := httpServer.Start(); err != nil {
			log.Error("Server failed to start: %v", err)
		}
	}()
	// Fiber, Handlers, Midleware:
	//app := handlers.InitFiber(srvcs, cfg)
	//
	//go func() {
	//	err := app.Listen(":8081")
	//	if err != nil {
	//		log.Error("server listening failed: %v", err)
	//	}
	//}()

	log.Info("server started")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	log.Info("Shutting down server...")
	if err := httpServer.Shutdown(); err != nil {
		log.Error("Server shutdown failed: %v", err)
	}
}
