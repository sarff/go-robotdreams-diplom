package services

import (
	"github.com/sarff/go-robotdreams-diplom/internal/clients"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
	"github.com/sarff/go-robotdreams-diplom/internal/repo"
)

type Services struct {
	Auth *AuthService
	Chat *ChatService
	WS   *WSService
}

func NewServices(cfg *config.Config, clnts *clients.Clients, repos *repo.Repos) (*Services, error) {
	return &Services{
		Auth: NewAuthService(cfg, clnts, repos),
		Chat: NewChatService(cfg, clnts, repos),
		WS:   NewWSService(cfg, clnts),
	}, nil

}
