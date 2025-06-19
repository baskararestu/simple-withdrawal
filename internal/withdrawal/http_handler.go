package withdrawal

import (
	"simple-withdraw-api/internal/domain"
	"simple-withdraw-api/internal/dto"
	"simple-withdraw-api/internal/middleware/validation"
	"simple-withdraw-api/internal/utilities"
	"simple-withdraw-api/internal/utilities/tools"

	"github.com/gofiber/fiber/v2"
)

type HttpWithdrawalHandler struct {
	withdrawalSvc domain.WithdrawalService
	secretKey string
}

func NewWithdrawalHandler(r fiber.Router, withdrawalSvc domain.WithdrawalService, secretKey string){
	handler := &HttpWithdrawalHandler{withdrawalSvc: withdrawalSvc, secretKey: secretKey}
	
	r.Post("/create",validation.New[dto.WriteTransactionRequestDto](),handler.CreateWithdrawHandler)
	r.Get("/", handler.GetAllHandler)
	r.Get("/:userId",handler.GetByUserIDHandler)
}

// CreateWithdrawHandler godoc
// @Summary Create a withdrawal
// @Description Create a withdrawal transaction for a user
// @Tags Withdrawal
// @Accept json
// @Produce json
// @Param X-TIMESTAMP header string true "Request timestamp"
// @Param X-SIGNATURE header string true "Request signature"
// @Param request body dto.WriteTransactionRequestDto true "Withdrawal Request"
// @Success 201 {object} map[string]interface{} "withdrawal successful"
// @Router /withdraw/create [post]
func (h *HttpWithdrawalHandler) CreateWithdrawHandler(c *fiber.Ctx) error {
	request := utilities.ExtractStructFromValidator[dto.WriteTransactionRequestDto](c)

	timestamp := c.Get("X-TIMESTAMP")
	signature := c.Get("X-SIGNATURE")

	if timestamp == "" || signature == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing X-TIMESTAMP or X-SIGNATURE headers")
	}

	method := c.Method()
	relativePath := c.Route().Path
	bodyBytes := c.Body()
	
	expectedSig, err := tools.GenerateSignature(method, relativePath, bodyBytes, timestamp, h.secretKey)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate signature")
	}

	if expectedSig != signature {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid signature")
	}

	err = h.withdrawalSvc.WriteTransaction(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "withdrawal successful",
	})
}

// GetByUserIDHandler godoc
// @Summary Get withdrawal history
// @Description Get withdrawal records by user ID
// @Tags Withdrawal
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {array} domain.Withdrawal
// @Router /withdraw/{userId} [get]
func (h *HttpWithdrawalHandler) GetByUserIDHandler(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("userId")
	if err != nil || userID <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	withdrawals, err := h.withdrawalSvc.GetByUserID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(withdrawals)
}

// GetAllHandler godoc
// @Summary Get all withdrawal records
// @Description Get all withdrawal transactions in the system
// @Tags Withdrawal
// @Accept json
// @Produce json
// @Success 200 {array} domain.Withdrawal
// @Router /withdraw [get]
func (h *HttpWithdrawalHandler) GetAllHandler(c *fiber.Ctx) error {
	withdrawals, err := h.withdrawalSvc.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(withdrawals)
}
