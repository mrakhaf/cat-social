package interfaces

import (
	"context"

	"github.com/mrakhaf/cat-social/models/dto"
	"github.com/mrakhaf/cat-social/models/request"
)

type Repository interface {
	SaveCat(ctx context.Context, req request.UploadCat, userId string) (data dto.SaveCatDto, err error)
}
