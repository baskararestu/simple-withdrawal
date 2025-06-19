package balance

import (
	"simple-withdraw-api/internal/domain"
	"simple-withdraw-api/internal/dto"

	"gorm.io/gorm"
)

type balanceService struct {
	balanceRepo domain.BalanceRepository
}

func NewBalanceService(balanceRepo domain.BalanceRepository) domain.BalanceService {
	return &balanceService{
		balanceRepo: balanceRepo,
	}
}

func (b *balanceService) GetByUserID(userID int) (domain.Balance, error) {
	return b.balanceRepo.GetByUserID(userID)
}

func (b *balanceService) UpdateAmount(userID int, amount float64) error {
	return b.balanceRepo.UpdateAmount(userID, amount)
}

func (b *balanceService) GenerateBalanceWithTx(tx *gorm.DB, req dto.GenerateBalanceRequestDto) error {
	balance := domain.Balance{
		UserID: req.UserID,
		Amount: req.Amount,
	}
	return b.balanceRepo.CreateWithTx(tx, balance)
}
