package handlers

import "github.com/sarff/go-robotdreams-diplom/internal/services"

type ChatHandler struct {
	chatService *services.ChatService
}

func NewChatHandler(chatService *services.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}
