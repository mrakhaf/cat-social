package interfaces

import (
	"context"

	"github.com/mrakhaf/cat-social/models/request"
)

type Usecase interface {
	SaveMatchCat(ctx context.Context, req request.MatchCat, userId string) (data interface{}, err error)
}
