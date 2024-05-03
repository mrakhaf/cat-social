package interfaces

import (
	"context"

	"github.com/mrakhaf/cat-social/models/dto"
	"github.com/mrakhaf/cat-social/models/entity"
	"github.com/mrakhaf/cat-social/models/request"
)

type Repository interface {
	SaveCat(ctx context.Context, req request.UploadCat, userId string) (data dto.SaveCatDto, err error)
	GetCats(ctx context.Context, query string) (cats []entity.Cat, err error)
	GetCatUser(ctx context.Context, userId, catId string) (cats entity.Cat, err error)
	UpdateCat(ctx context.Context, catId, userId string, req request.UploadCat) (err error)
	DeleteCat(ctx context.Context, catId string) (err error)
	GetCatUserHasNotMatch(ctx context.Context, userId, catId string) (sex string, err error)
	GetCatMatchHasNotMatch(ctx context.Context, userId, catId string) (sex string, err error)
	SaveMatchCat(ctx context.Context, userId string, req request.MatchCat) (matchId string, err error)
}
