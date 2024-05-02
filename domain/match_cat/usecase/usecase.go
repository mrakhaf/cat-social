package usecase

import (
	"context"
	"fmt"

	"github.com/mrakhaf/cat-social/domain/match_cat/interfaces"
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

func (u *usecase) SaveMatchCat(ctx context.Context, req request.MatchCat, userId string) (data interface{}, err error) {

	data, err = u.repository.SaveMatchCat(ctx, req, userId)

	if err != nil {
		err = fmt.Errorf("failed to save match cat: %s", err)
		return
	}

	return

}
