package repository

import (
	"context"
	"database/sql"

	"github.com/mrakhaf/cat-social/domain/auth/interfaces"
)

type repoHandler struct {
	catDB *sql.DB
}

func NewRepository(catDB *sql.DB) interfaces.Repository {
	return &repoHandler{
		catDB: catDB,
	}
}

func (r *repoHandler) Login(ctx context.Context, email, password string) (string, error) {
	// TODO
	return "", nil
}
