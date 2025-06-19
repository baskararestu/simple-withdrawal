package balance

import (
	"simple-withdraw-api/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type HttpBalanceHandler struct {
	balanceSvc domain.BalanceService
	secretKey  string
}

func NewBalanceHttpHandler(r fiber.Router, balanceSvc domain.BalanceService, secretKey string) {
	handler := &HttpBalanceHandler{
		balanceSvc: balanceSvc,
		secretKey:  secretKey,
	}

	r.Get("/inquiry/:userId", handler.InquiryUserBalance)
	r.Get("/inquiry", handler.InquiryBalance)
}

// InquiryUserBalance godoc
// @Summary      Inquiry Balance by User
// @Description  Check user balance by user ID using secret key header
// @Tags         Balance
// @Accept       json
// @Produce      json
// @Param        userId path int true "User ID"
// @Param        X-SECRET-KEY header string true "Secret Key" example(dev-secret)
// @Success      200 {object} domain.Balance
// @Router       /balance/inquiry/{userId} [get]
func (h *HttpBalanceHandler) InquiryUserBalance(c *fiber.Ctx) error {
	secretKey := c.Get("X-SECRET-KEY")
	if secretKey == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing X-SECRET-KEY header")
	}

	if secretKey != h.secretKey {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid secret key")
	}

	userID, err := c.ParamsInt("userId")
	if err != nil || userID <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	balance, err := h.balanceSvc.GetByUserID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(balance)
}

// InquiryBalance godoc
// @Summary      Inquiry All Balances
// @Description  Retrieve all balance records
// @Tags         Balance
// @Accept       json
// @Produce      json
// @Param        X-SECRET-KEY header string true "Secret Key" example(dev-secret)
// @Success      200 {array} domain.Balance
// @Router       /balance/inquiry [get]
func (h *HttpBalanceHandler) InquiryBalance(c *fiber.Ctx) error {
	secretKey := c.Get("X-SECRET-KEY")
	if secretKey == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing X-SECRET-KEY header")
	}

	if secretKey != h.secretKey {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid secret key")
	}

	balances, err := h.balanceSvc.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(balances)
}
