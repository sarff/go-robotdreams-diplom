package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sarff/go-robotdreams-diplom/internal/models"
	"github.com/sarff/go-robotdreams-diplom/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// @Summary      Реєстрація нового користувача
// @Description  Створення облікового запису користувача
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      models.RegisterRequest  true  "Дані для реєстрації"
// @Success      200      {object}  map[string]string        "Статус успішної реєстрації"
// @Failure      400      {object}  map[string]string        "Невірні дані"
// @Router       /api/v1/auth/register [post]
func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.authService.Register(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
	})
}

// Login godoc
// @Summary      Вхід користувача
// @Description  Аутентифікація користувача та отримання токена
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      models.LoginRequest  true  "Дані для входу"
// @Success      200      {object}  map[string]interface{} "Користувач та токен"
// @Failure      401      {object}  map[string]string       "Невірний логін або пароль"
// @Router       /api/v1/auth/login [post]
func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, token, err := h.authService.Login(&req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"user":  user,
		"token": token,
	})
}

//func (h *AuthHandler) GetProfile(c fiber.Ctx) error {
//	userID := c.Locals("userID").(string)
//
//	user, err := h.authService.GetUserByID(userID)
//	if err != nil {
//		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//			"error": "User not found",
//		})
//	}
//
//	return c.JSON(user)
//}
