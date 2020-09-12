package domain

import (
	"context"
	"time"
)

// User struct
type User struct {
	ID        int        `json:"user_id" gorm:"primary_key"`
	Name      string     `json:"name" validate:"required"`
	Email     string     `json:"email" validate:"required,email"`
	Password  string     `json:"password" validate:"min=6,max=75"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// UserUsecase interface
type UserUsecase interface {
	GetByID(etx context.Context, id int) (User, error)
	Store(etx context.Context, u *User) error
	Update(etx context.Context, t *User) error
	Delete(etx context.Context, id int) error
}

// UserRepository interface
type UserRepository interface {
	GetByID(ctx context.Context, id int) (User, error)
	Store(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	Delete(etx context.Context, id int) error
}
