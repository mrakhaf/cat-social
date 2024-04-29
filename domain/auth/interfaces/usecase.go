package interfaces

import (
	"context"

	"github.com/mrakhaf/cat-social/models/dto"
)

type Usecase interface {
	Login(ctx context.Context, email, password string) (data dto.Login, err error)
}
