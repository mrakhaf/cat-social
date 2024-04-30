package usecase

import (
	"context"
	"fmt"

	"github.com/mrakhaf/cat-social/domain/cat/interfaces"
	"github.com/mrakhaf/cat-social/models/request"
)

type usecase struct {
	repository interfaces.Repository
}

func NewUsecase(repository interfaces.Repository) interfaces.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) UploadCat(ctx context.Context, req request.UploadCat, userId string) (data interface{}, err error) {

	data, err = u.repository.SaveCat(ctx, req, userId)

	if err != nil {
		err = fmt.Errorf("failed to save cat: %s", err)
		return
	}

	return

}
