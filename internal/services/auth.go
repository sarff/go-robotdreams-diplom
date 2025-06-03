package services

import (
	"github.com/sarff/go-robotdreams-diplom/internal/clients"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
	"github.com/sarff/go-robotdreams-diplom/internal/models"
)

type AuthService struct {
	cfg   *config.Config
	clnts *clients.Clients
}

func NewAuthService(cfg *config.Config, clnts *clients.Clients) *AuthService {
	return &AuthService{
		cfg:   cfg,
		clnts: clnts,
	}
}

func (as *AuthService) Register(req *models.RegisterRequest) error {
	return nil
}

func (as *AuthService) Login(req *models.LoginRequest) error {
	return nil
}

func (as *AuthService) Logout() error {
	return nil
}
