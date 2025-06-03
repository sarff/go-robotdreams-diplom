package services

import (
	"github.com/sarff/go-robotdreams-diplom/internal/clients"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
)

type WSService struct {
	cfg   *config.Config
	clnts *clients.Clients
}

func NewWSService(cfg *config.Config, clnts *clients.Clients) *WSService {
	return &WSService{
		cfg:   cfg,
		clnts: clnts,
	}
}
