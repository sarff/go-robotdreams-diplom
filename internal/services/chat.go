package services

import (
	"github.com/sarff/go-robotdreams-diplom/internal/clients"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
)

type ChatService struct {
	cfg   *config.Config
	clnts *clients.Clients
}

func NewChatService(cfg *config.Config, clnts *clients.Clients) *ChatService {
	return &ChatService{
		cfg:   cfg,
		clnts: clnts,
	}
}
