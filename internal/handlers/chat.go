package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sarff/go-robotdreams-diplom/internal/models"
	"github.com/sarff/go-robotdreams-diplom/internal/services"
)

type ChatHandler struct {
	chatService *services.ChatService
}

func NewChatHandler(chatService *services.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}

// @Summary      Відправити повідомлення
// @Description  Відправка повідомлення в канал/room
// @Tags         chat
// @Accept       json
// @Produce      json
// @Security     UserTokenAuth
// @Param        request  body      models.MessageRequest  true  "Дані для відправки повідомлення"
// @Success      201      {object}  models.MessageResponse  "Канал і повідомлення" // якщо є така структура
// @Failure      400      {object}  map[string]string       "Помилка валідації або бізнес-логіки"
// @Failure      401      {object}  map[string]string       "Неавторизований доступ"
// @Router       /api/v1/chat/messages [post]
func (ch *ChatHandler) SendMessage(c fiber.Ctx) error {
	var req models.MessageRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	userID := c.Locals("userID").(string)

	message, err := ch.chatService.SendMessage(userID, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(message)
}

func (ch *ChatHandler) GetRooms(c fiber.Ctx) error {
	// TODO: need implement
	return nil
}
func (ch *ChatHandler) CreateRoom(c fiber.Ctx) error {
	// TODO: need implement
	return nil
}
