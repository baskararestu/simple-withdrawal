package user

import (
	"simple-withdraw-api/internal/domain"
	"simple-withdraw-api/internal/dto"
	"simple-withdraw-api/internal/middleware/validation"
	"simple-withdraw-api/internal/utilities"

	"github.com/gofiber/fiber/v2"
)

type HttpUserHandler struct{
	userSvc domain.UserService
}

func NewUserHttpHandler(r fiber.Router, userSvc domain.UserService) {
	handler := &HttpUserHandler{
		userSvc: userSvc,
	}

	r.Post("/create",validation.New[dto.CreateUserRequestDto](), handler.CreateUserHandler)
	r.Get("/",handler.GetUserList)
	r.Get("/:userId",handler.GetUserByID)
}

// CreateUserHandler godoc
// @Summary Create new user
// @Description Create a new user and generate initial balance
// @Tags User
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequestDto true "Create User Request"
// @Success 201 {object} map[string]interface{}
// @Router /user/create [post]
func (h *HttpUserHandler) CreateUserHandler(c *fiber.Ctx) error {
	request := utilities.ExtractStructFromValidator[dto.CreateUserRequestDto](c)

	if request == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	err := h.userSvc.CreateUser(*request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user successfully created",
	})
}

// GetUserList godoc
// @Summary Get all users
// @Description Get list of all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} domain.User
// @Router /user [get]
func (h *HttpUserHandler) GetUserList(c *fiber.Ctx) error {
	users, err := h.userSvc.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(users)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get a specific user by user ID
// @Tags User
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} domain.User
// @Router /user/{userId} [get]
func (h *HttpUserHandler) GetUserByID(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("userId")
	if err != nil || userID <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	user, err := h.userSvc.GetByID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(user)
}
