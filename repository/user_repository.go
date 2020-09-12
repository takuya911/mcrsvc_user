package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/takuya911/mcrsvc_user/domain"
)

type userRepo struct {
	Conn *gorm.DB
}

// NewUserRepository function
func NewUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &userRepo{Conn}
}

func (u *userRepo) GetByID(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	if result := u.Conn.Where("id = ? AND deleted_at is null", id).Find(&user); result.Error != nil {
		return domain.User{}, result.Error
	}
	return user, nil
}

func (u *userRepo) Store(ctx context.Context, form *domain.User) (err error) {
	if result := u.Conn.Create(&form); result.Error != nil {
		return result.Error
	}
	return
}

func (u *userRepo) Update(ctx context.Context, form *domain.User) (err error) {
	if result := u.Conn.Model(&form).Where("id = ?", form.ID).Updates(&form); result.Error != nil {
		return result.Error
	}
	return
}

func (u *userRepo) Delete(ctx context.Context, id int) (err error) {
	if result := u.Conn.Where("id = ?", id).Delete(&domain.User{}); result.Error != nil {
		return result.Error
	}
	return
}
