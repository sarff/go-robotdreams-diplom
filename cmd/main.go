package main

import (
	"context"
	"fmt"

	"github.com/sarff/go-robotdreams-diplom/internal/clients"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
	log "github.com/sarff/iSlogger"
)

func main() {
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

	// Clients
	cl, err := clients.NewClients(ctx, cfg)
	if err != nil {
		log.Error("Failed to create clients: %v", err)
		panic(err)
	}

	// Services TODO:

	// Middlewares TODO:

	// Handlers TODO:
}
