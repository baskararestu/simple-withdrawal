package user

import (
	"simple-withdraw-api/internal/domain"
	"simple-withdraw-api/internal/dto"

	"gorm.io/gorm"
)

type userService struct {
	userRepo domain.UserRepository
	balanceSvc domain.BalanceService
}

func NewUserService(userRepo domain.UserRepository, balanceSvc domain.BalanceService) domain.UserService {
	return &userService{
		userRepo: userRepo,
		balanceSvc: balanceSvc,
	}
}

func (s *userService) GetAll() ([]domain.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) GetByID(id int) (*domain.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) CreateUser(request dto.CreateUserRequestDto) error {
	return s.userRepo.WithTransaction(func(tx *gorm.DB) error {
		newUser := domain.User{Name: request.Name}

		// Simpan user dan rollback kalau gagal
		if err := s.userRepo.CreateWithTx(tx, &newUser); err != nil {
			return err
		}

		// Simpan balance
		balance := dto.GenerateBalanceRequestDto{
			UserID: newUser.ID,
			Amount: request.Amount,
		}
		if err := s.balanceSvc.GenerateBalanceWithTx(tx, balance); err != nil {
			return err
		}

		return nil
	})
}



