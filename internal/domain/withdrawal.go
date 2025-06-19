package domain

import (
	"simple-withdraw-api/internal/dto"
	"time"
)

type Withdrawal struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `json:"user_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type WithdrawalRepository interface {
	Create(withdrawal Withdrawal) error
	FindByUserID(userID int) ([]Withdrawal, error)
	FindAll() ([]Withdrawal,error)
}

type WithdrawalService interface {
	WriteTransaction(req *dto.WriteTransactionRequestDto) error
	GetByUserID(userID int) ([]Withdrawal, error)
	GetAll() ([]Withdrawal,error)
}