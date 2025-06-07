package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sarff/go-robotdreams-diplom/internal/models"
	"github.com/sarff/go-robotdreams-diplom/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

var validate = validator.New()

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

	// Validate check
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
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

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
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

// @Summary      Отримати профіль користувача
// @Description  Отримати інформацію про поточного автентифікованого користувача
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     UserTokenAuth
// @Success      200  {object}  models.User              "Профіль користувача"
// @Failure      401  {object}  map[string]string        "Неавторизований доступ"
// @Failure      404  {object}  map[string]string        "Користувача не знайдено"
// @Router       /api/v1/auth/profile [get]
func (h *AuthHandler) GetProfile(c fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	user, err := h.authService.FindByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}
