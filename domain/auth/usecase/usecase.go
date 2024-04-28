package usecase

import (
	"context"

	"github.com/mrakhaf/cat-social/domain/auth/interfaces"
)

type (
	usecase struct {
		repository interfaces.Repository
	}
)

func NewUsecase(repository interfaces.Repository) interfaces.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) Login(ctx context.Context, email, password string) (string, error) {
	return u.repository.Login(ctx, email, password)
}
