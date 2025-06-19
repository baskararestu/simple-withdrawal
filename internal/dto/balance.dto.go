package dto

type GenerateBalanceRequestDto struct{
	UserID int `json:"userId"`
	Amount float64 `json:"amount"`
}

type BalanceInquiryRequestDto struct {
	UserID    int    `json:"userId" validate:"required" example:"1"`
	SecretKey string `json:"secretKey" validate:"required" example:"dev-secret"`
}
