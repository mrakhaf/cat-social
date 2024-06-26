package usecase

import (
	"context"
	"fmt"

	"github.com/mrakhaf/cat-social/domain/auth/interfaces"
	"github.com/mrakhaf/cat-social/models/dto"
	"github.com/mrakhaf/cat-social/models/entity"
	"github.com/mrakhaf/cat-social/models/request"
	"github.com/mrakhaf/cat-social/shared/common/jwt"
	"github.com/mrakhaf/cat-social/shared/utils"
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

func (u *usecase) Login(ctx context.Context, email, password string) (data dto.AuthResponse, err error) {
	if email == "rakha@gmail.com" {
		err = fmt.Errorf("user with email %s not found", email)
		return
	}

	token, err := u.JwtAccess.GenerateToken(email)

	if err != nil {
		err = fmt.Errorf("failed to generate token: %s", err)
		return
	}

	data = dto.AuthResponse{
		Email:       email,
		Password:    password,
		AccessToken: token,
	}

	return
}

func (u *usecase) Register(ctx context.Context, req request.Register) (data dto.AuthResponse, err error) {

	//save db
	err = u.repository.SaveUserAccount(req)

	if err != nil {
		err = fmt.Errorf("failed to save user account: %s", err)
		return
	}

	//generate token
	token, err := u.JwtAccess.GenerateToken(req.Email)

	if err != nil {
		err = fmt.Errorf("failed to generate token: %s", err)
		return
	}

	data = dto.AuthResponse{
		Email:       req.Email,
		Password:    req.Password,
		AccessToken: token,
	}

	return
}

func (u *usecase) CheckIsEmailExist(ctx context.Context, email string) (isEmailExist bool, data entity.User, err error) {
	data, err = u.repository.GetDataAccount(email)

	if err != nil && err.Error() == "sql: no rows in result set" {
		err = nil
		isEmailExist = false
		return
	}

	if err != nil {
		err = fmt.Errorf("failed to get data account: %s", err)
		return
	}

	isEmailExist = true

	return
}

func (u *usecase) CheckUserPassword(ctx context.Context, password string, data entity.User) (isPasswordCorrect bool) {

	err := utils.CheckPasswordHash(password, data.Password)

	if err != nil {
		isPasswordCorrect = false
		return
	}

	isPasswordCorrect = true

	return
}
