package usecase

import (
	"context"
	"fmt"

	"github.com/mrakhaf/cat-social/domain/auth/interfaces"
	"github.com/mrakhaf/cat-social/models/dto"
	"github.com/mrakhaf/cat-social/shared/common/jwt"
)

type (
	usecase struct {
		repository interfaces.Repository
		JwtAccess  *jwt.JWT
	}
)

func NewUsecase(repository interfaces.Repository, JwtAccess *jwt.JWT) interfaces.Usecase {
	return &usecase{
		repository: repository,
		JwtAccess:  JwtAccess,
	}
}

func (u *usecase) Login(ctx context.Context, email, password string) (data dto.Login, err error) {
	if email == "rakha@gmail.com" {
		err = fmt.Errorf("user with email %s not found", email)
		return
	}

	token, err := u.JwtAccess.GenerateToken(email)

	if err != nil {
		err = fmt.Errorf("failed to generate token: %s", err)
		return
	}

	data = dto.Login{
		Email:       email,
		Password:    password,
		AccessToken: token,
	}

	return
}
