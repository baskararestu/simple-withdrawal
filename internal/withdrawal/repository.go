package withdrawal

import (
	"simple-withdraw-api/internal/domain"

	"gorm.io/gorm"
)

type withdrawalRepository struct {
	db *gorm.DB
}

func NewWithdrawalRepository(db *gorm.DB) domain.WithdrawalRepository {
	return &withdrawalRepository{db}
}

func (r *withdrawalRepository) Create(withdrawal domain.Withdrawal) error {
	return r.db.Create(&withdrawal).Error
}

func (r *withdrawalRepository) FindByUserID(userID int) ([]domain.Withdrawal, error) {
	var withdrawals []domain.Withdrawal
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&withdrawals).Error
	return withdrawals, err
}

func (r *withdrawalRepository) FindAll() ([]domain.Withdrawal, error) {
	var withdrawals []domain.Withdrawal
	err := r.db.Order("created_at DESC").Find(&withdrawals).Error
	return withdrawals, err
}

