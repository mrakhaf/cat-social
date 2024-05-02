package interfaces

import (
	"context"

	"github.com/mrakhaf/cat-social/models/dto"
	"github.com/mrakhaf/cat-social/models/request"
)

type Repository interface {
	SaveMatchCat(ctx context.Context, req request.MatchCat, userId string) (data dto.SaveMatchCatDto, err error)
}
