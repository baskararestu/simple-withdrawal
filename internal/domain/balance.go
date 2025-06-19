package domain

import (
	"simple-withdraw-api/internal/dto"
	"time"

	"gorm.io/gorm"
)

type Balance struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `json:"user_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type BalanceRepository interface {
	CreateWithTx(tx *gorm.DB, balance Balance) error
	GetByUserID(userID int) (Balance, error)
	UpdateAmount(userID int, amount float64) error
	FindAll() ([]Balance,error)
}

type BalanceService interface {
	GenerateBalanceWithTx(tx *gorm.DB, req dto.GenerateBalanceRequestDto) error 	
	GetByUserID(userID int) (Balance,error)
	UpdateAmount(userID int, amount float64) error
	GetAll() ([]Balance,error)
}