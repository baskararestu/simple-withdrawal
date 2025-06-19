package user

import (
	"simple-withdraw-api/internal/domain"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) domain.UserRepository {
	return &userRepo{db}
}

func (r *userRepo) FindAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) FindByID(id int) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) CreateWithTx(tx *gorm.DB, user *domain.User) error {
	return tx.Create(user).Error
}

func (r *userRepo) WithTransaction(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

