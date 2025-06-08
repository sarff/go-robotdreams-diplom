package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sarff/go-robotdreams-diplom/internal/services"
)

type WebsocketHandler struct {
	ws *services.WSService
	ch *services.ChatService
}

func NewWSHandler(ws *services.WSService, ch *services.ChatService) *WebsocketHandler {
	return &WebsocketHandler{
		ws: ws,
		ch: ch,
	}
}

func (ch *WebsocketHandler) HandleWebSocket(c fiber.Ctx) error {
	// TODO:
	return nil
}
