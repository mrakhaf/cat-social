package interfaces

import (
	"context"

	"github.com/mrakhaf/cat-social/models/dto"
	"github.com/mrakhaf/cat-social/models/entity"
	"github.com/mrakhaf/cat-social/models/request"
)

type Usecase interface {
	Login(ctx context.Context, email, password string) (data dto.AuthResponse, err error)
	Register(ctx context.Context, req request.Register) (data dto.AuthResponse, err error)
	CheckIsEmailExist(ctx context.Context, email string) (isEmailExist bool, data entity.User, err error)
	CheckUserPassword(ctx context.Context, password string, data entity.User) (isPasswordCorrect bool)
}
