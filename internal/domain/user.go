package domain

import (
	"simple-withdraw-api/internal/dto"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id int) (*User, error)
	CreateWithTx(tx *gorm.DB, user *User) error
	WithTransaction(fn func(tx *gorm.DB) error) error
}

type UserService interface {
	GetAll() ([]User,error)
	GetByID(id int) (*User,error)
	CreateUser(request dto.CreateUserRequestDto) error
}
