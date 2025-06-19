package dto

type CreateUserRequestDto struct {
	Name string `json:"name"`
	Amount float64 `json:"amount"`
}