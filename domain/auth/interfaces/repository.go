package interfaces

import (
	"context"

	"github.com/mrakhaf/cat-social/models/entity"
	"github.com/mrakhaf/cat-social/models/request"
)

type Repository interface {
	Login(ctx context.Context, email, password string) (string, error)
	SaveUserAccount(data request.Register) (err error)
	GetDataAccount(email string) (data entity.User, err error)
}
