package interfaces

import (
	"context"

	"github.com/mrakhaf/cat-social/models/request"
)

type Usecase interface {
	UploadCat(ctx context.Context, req request.UploadCat, userId string) (data interface{}, err error)
	GetCat(ctx context.Context, userId string, req request.GetCatParam) (data interface{}, err error)
	ValidationCatUser(ctx context.Context, catId, userId string) (err error)
	UpdateCat(ctx context.Context, catId, userId string, req request.UploadCat) (err error)
	DeleteCat(ctx context.Context, catId string) (err error)
	ValidateMatchCat(ctx context.Context, userId string, req request.MatchCat) (err error)
	UploadMatch(ctx context.Context, req request.MatchCat, userId string) (err error)
}
