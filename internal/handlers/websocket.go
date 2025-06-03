package handlers

import "github.com/sarff/go-robotdreams-diplom/internal/services"

type WebsocketHandler struct {
	wsService *services.WSService
}

func NewWSHandler(wsService *services.WSService) *WebsocketHandler {
	return &WebsocketHandler{wsService: wsService}
}
