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
}
