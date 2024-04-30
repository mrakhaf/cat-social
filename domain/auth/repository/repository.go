package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mrakhaf/cat-social/domain/auth/interfaces"
	"github.com/mrakhaf/cat-social/models/entity"
	"github.com/mrakhaf/cat-social/models/request"
	"github.com/mrakhaf/cat-social/shared/utils"
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

func (r *repoHandler) SaveUserAccount(data request.Register) (err error) {
	id := utils.GenerateUUID()
	fmt.Println(id)

	hasPassword, _ := utils.HashPassword(data.Password)

	query := fmt.Sprintf(`INSERT INTO users (id, email, password, name) VALUES ('%s', '%s', '%s', '%s')`, id, data.Email, hasPassword, data.Name)

	_, err = r.catDB.Exec(query)

	if err != nil {
		return
	}

	return
}

func (r *repoHandler) GetDataAccount(email string) (data entity.User, err error) {

	row := r.catDB.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", email))

	err = row.Scan(&data.Id, &data.Name, &data.Email, &data.Password)

	if err != nil {
		return
	}

	return
}
