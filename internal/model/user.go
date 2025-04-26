package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID            string     `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	Email           string     `json:"email,omitempty" gorm:"unique"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Password        string     `json:"-"`
	RememberToken   string     `json:"remember_token,omitempty"`
	Status          string     `json:"status,omitempty"`
}

func (User) TableName() string {
	return "users"
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindByUuid(ctx context.Context, uuid string) (User, error)
	FindById(ctx context.Context, id uint) (User, error)
	IsExist(ctx context.Context, email string) error
	Save(ctx context.Context, c *User) error
	Update(ctx context.Context, c *User) error
	Delete(ctx context.Context, id uint) error
}

type UserService interface {
	Index(ctx context.Context) ([]User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
	SaveUser(ctx context.Context, user *User) error
	IsUserExist(ctx context.Context, email string) error
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Firstname string `json:"first_name" validate:"required,min=3"`
	Lastname  string `json:"last_name"`
}
