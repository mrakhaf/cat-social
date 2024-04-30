package interfaces

import (
	"context"

	"github.com/mrakhaf/cat-social/models/request"
)

type Usecase interface {
	UploadCat(ctx context.Context, req request.UploadCat, userId string) (data interface{}, err error)
}
