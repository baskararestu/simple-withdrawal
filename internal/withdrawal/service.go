package withdrawal

import (
	"errors"
	"simple-withdraw-api/internal/domain"
	"simple-withdraw-api/internal/dto"
)

type withdrawalService struct {
	userSvc      domain.UserService
	balanceSvc   domain.BalanceService
	withdrawRepo domain.WithdrawalRepository
}

func NewWithdrawalService(
	userSvc domain.UserService,
	balanceSvc domain.BalanceService,
	withdrawRepo domain.WithdrawalRepository,
) domain.WithdrawalService {
	return &withdrawalService{
		userSvc:      userSvc,
		balanceSvc:   balanceSvc,
		withdrawRepo: withdrawRepo,
	}
}

func (w *withdrawalService) GetByUserID(userID int) ([]domain.Withdrawal, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	_, err := w.userSvc.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return w.withdrawRepo.FindByUserID(userID)
}

func (w *withdrawalService) WriteTransaction(req *dto.WriteTransactionRequestDto) error {
	userID := req.UserID
	amount := req.Amount

	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than 0")
	}

	user, err := w.userSvc.GetByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	balance, err := w.balanceSvc.GetByUserID(user.ID)
	if err != nil {
		return errors.New("balance not found")
	}

	if balance.Amount < amount {
		return errors.New("insufficient balance")
	}

	newBalance := balance.Amount - amount
	if err := w.balanceSvc.UpdateAmount(user.ID, newBalance); err != nil {
		return err
	}

	withdraw := domain.Withdrawal{
		UserID: user.ID,
		Amount: amount,
	}
	return w.withdrawRepo.Create(withdraw)
}

func (w *withdrawalService) GetAll() ([]domain.Withdrawal, error) {
	return w.withdrawRepo.FindAll()
}
