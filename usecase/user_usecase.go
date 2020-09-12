package usecase

import (
	"context"

	"github.com/takuya911/mcrsvc_user/conf"
	"github.com/takuya911/mcrsvc_user/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

// NewUserUsecase func
func NewUserUsecase(u domain.UserRepository) *userUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) GetByID(ctx context.Context, id int) (domain.User, error) {
	result, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return result, nil
}

func (u *userUsecase) Store(ctx context.Context, form *domain.User) (err error) {
	err = u.userRepo.Store(ctx, form)
	return
}

func (u *userUsecase) Update(ctx context.Context, form *domain.User) (err error) {
	if _, err := u.userRepo.GetByID(ctx, form.ID); err != nil {
		return conf.ErrNotExist
	}
	return u.userRepo.Update(ctx, form)
}

func (u *userUsecase) Delete(ctx context.Context, id int) (err error) {
	if _, err := u.userRepo.GetByID(ctx, id); err != nil {
		return conf.ErrNotExist
	}
	return u.userRepo.Delete(ctx, id)
}
