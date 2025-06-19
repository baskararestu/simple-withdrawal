package tools

import (
	"encoding/json"
	"simple-withdraw-api/internal/dto"
	"simple-withdraw-api/internal/middleware/validation"
	"simple-withdraw-api/internal/utilities"

	"github.com/gofiber/fiber/v2"
)

// HttpToolsHandler handles tools-related endpoints.
type HttpToolsHandler struct{
	secretKey string
}

// NewToolsHttpHandler registers tools routes.
func NewToolsHttpHandler(r fiber.Router, secretKey string) {
	handler := &HttpToolsHandler{secretKey: secretKey}
	r.Post("/signature", validation.New[dto.SignatureRequestDto](), handler.SignatureGeneratorHandler)
}

// SignatureGeneratorHandler generates a signature from request data.
// @Summary      Generate Signature
// @Description  Generate SHA-256 signature based on method, relative path, body, and timestamp.
// @Tags         Tools
// @Accept       json
// @Produce      json
// @Param request body dto.SignatureRequestExample true "Signature Request Payload (param is request body for protected service)"
// @Success 200 {object} dto.SignatureResponseDto
// @Router       /tools/signature [post]
func (h *HttpToolsHandler) SignatureGeneratorHandler(c *fiber.Ctx) error {
	request := utilities.ExtractStructFromValidator[dto.SignatureRequestDto](c)
	bodyBytes,err := json.Marshal(request.Param)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	signature, err := GenerateSignature(
		request.Method,
		request.RelativePath,
		bodyBytes,
		request.Timestamp,
		h.secretKey,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"signature": signature,
	})
}
