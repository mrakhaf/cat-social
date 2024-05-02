package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mrakhaf/cat-social/domain/match_cat/interfaces"
	"github.com/mrakhaf/cat-social/models/dto"
	"github.com/mrakhaf/cat-social/models/request"
	"github.com/mrakhaf/cat-social/shared/utils"
)

type repoHandler struct {
	matchCatDB *sql.DB
}

func NewRepository(matchCatDB *sql.DB) interfaces.Repository {
	return &repoHandler{
		matchCatDB: matchCatDB,
	}
}

func (r *repoHandler) SaveMatchCat(ctx context.Context, req request.MatchCat, userId string) (data dto.SaveMatchCatDto, err error) {

	idMatchCat := utils.GenerateUUID()
	createdAt := utils.GenerateDatetimeIso8601()

	querySaveMatchCat := fmt.Sprintf(`INSERT INTO match_cats (id, matchcatId, userCatId, createdAt) VALUES ('%s', '%s', '%s', '%s')`, idMatchCat, req.MatchCatId, req.UserCatId, createdAt)

	_, err = r.matchCatDB.Exec(querySaveMatchCat)

	if err != nil {
		err = fmt.Errorf("failed to save match cat: %s", err)
		return
	}

	data = dto.SaveMatchCatDto{
		Id:        idMatchCat,
		CreatedAt: createdAt,
	}

	return
}
