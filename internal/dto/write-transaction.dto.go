package dto

type WriteTransactionRequestDto struct {
	UserID int `json:"userId"`
	Amount float64 `json:"amount"`
}
