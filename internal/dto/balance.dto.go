package dto

type GenerateBalanceRequestDto struct{
	UserID int `json:"userId"`
	Amount float64 `json:"amount"`
}