package services

import (
	"github.com/sarff/go-robotdreams-diplom/internal/clients"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
)

type Services struct {
	Auth *AuthService
	Chat *ChatService
	WS   *WSService
}

func NewServices(cfg *config.Config, clnts *clients.Clients) (*Services, error) {
	return &Services{
		Auth: NewAuthService(cfg, clnts),
		Chat: NewChatService(cfg, clnts),
		WS:   NewWSService(cfg, clnts),
	}, nil

}
