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

// @Summary      Отримати всі кімнати користувача
// @Description  Повертає список кімнат, в яких присутній авторизований користувач
// @Tags         chat
// @Accept       json
// @Produce      json
// @Security     UserTokenAuth
// @Success      200  {array}   models.Room           "Список кімнат"
// @Failure      400  {object}  map[string]string      "Не вдалося отримати кімнати"
// @Failure      401  {object}  map[string]string      "Неавторизований доступ"
// @Router       /api/v1/chat/rooms [get]
func (ch *ChatHandler) GetRooms(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	room, err := ch.chatService.GetUserRooms(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot receive rooms for user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(room)
}

// @Summary      Створення кімнати
// @Description Створення кімнати для спілкування
// @Tags         chat
// @Accept       json
// @Produce      json
// @Security     UserTokenAuth
// @Param        request  body      models.CreateRoomRequest  true  "Дані для створення"
// @Failure      400      {object}  map[string]string       "Помилка валідації або бізнес-логіки"
// @Failure      401      {object}  map[string]string       "Неавторизований доступ"
// @Router       /api/v1/chat/rooms [post]
func (ch *ChatHandler) CreateRoom(c fiber.Ctx) error {
	var req models.CreateRoomRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	// Validate check
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID := c.Locals("userID").(string)

	room, err := ch.chatService.CreateRoom(userID, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(room)
}

// @Summary      Отримати кімнату по ID
// @Description  Отримати інформацію про кімнату за її ID
// @Tags         chat
// @Accept       json
// @Produce      json
// @Security     UserTokenAuth
// @Param        roomID  path      string  true  "ID кімнати"
// @Success      200  {object}  models.Room              "Кімната"
// @Failure      401  {object}  map[string]string        "Неавторизований доступ"
// @Failure      404  {object}  map[string]string        "Кімната не знайдена"
// @Router       /api/v1/chat/id/{roomID} [get]
func (ch *ChatHandler) FindRoomByID(c fiber.Ctx) error {
	roomIDParam := c.Params("roomID")

	room, err := ch.chatService.FindRoomByID(roomIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}
	return c.Status(fiber.StatusOK).JSON(room)
}

// @Summary      Отримати кімнату по Name
// @Description  Отримати інформацію про кімнату за її Name
// @Tags         chat
// @Accept       json
// @Produce      json
// @Security     UserTokenAuth
// @Param        roomName  path      string  true  "Name кімнати"
// @Success      200  {object}  models.Room              "Кімната"
// @Failure      401  {object}  map[string]string        "Неавторизований доступ"
// @Failure      404  {object}  map[string]string        "Кімната не знайдена"
// @Router       /api/v1/chat/rooms/{roomName} [get]
func (ch *ChatHandler) FindRoomByName(c fiber.Ctx) error {
	roomNameParam := c.Params("roomName")

	room, err := ch.chatService.FindRoomByName(roomNameParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}
	return c.Status(fiber.StatusOK).JSON(room)
}

func (ch *ChatHandler) GetMessages(ctx fiber.Ctx) error {
	//	TODO
	return nil
}
