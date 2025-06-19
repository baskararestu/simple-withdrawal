package balance

import (
	"simple-withdraw-api/internal/domain"

	"gorm.io/gorm"
)

type balanceRepository struct {
	db *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) domain.BalanceRepository {
	return &balanceRepository{db}
}

func (r *balanceRepository) CreateWithTx(tx *gorm.DB, balance domain.Balance) error {
	return tx.Create(&balance).Error
}

func (r *balanceRepository) GetByUserID(userID int) (domain.Balance, error) {
	var balance domain.Balance
	err := r.db.Where("user_id = ?", userID).First(&balance).Error
	return balance, err
}

func (r *balanceRepository) UpdateAmount(userID int, amount float64) error {
	return r.db.Model(&domain.Balance{}).
		Where("user_id = ?", userID).
		Update("amount", amount).Error
}
