package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mrakhaf/cat-social/domain/cat/interfaces"
	"github.com/mrakhaf/cat-social/models/dto"
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

func (r *repoHandler) SaveCat(ctx context.Context, req request.UploadCat, userId string) (data dto.SaveCatDto, err error) {

	idCat := utils.GenerateUUID()
	createdAt := time.Now()

	imageUrls := strings.Join(req.ImageUrls, ",")

	querySaveCat := fmt.Sprintf(`INSERT INTO cats (id, name, ageInMonths, race, sex, description, userId, imageUrls, createdAt) VALUES ('%s', '%s', %d, '%s', '%s', '%s', '%s', '%s', '%s')`, idCat, req.Name, req.AgeInMonth, req.Race, req.Sex, req.Description, userId, imageUrls, createdAt.Format(time.RFC3339))

	tx, err := r.catDB.Begin()

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = tx.Exec(querySaveCat)

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
		return
	}

	data = dto.SaveCatDto{
		Id:        idCat,
		CreatedAt: utils.ConvertDateIso(createdAt),
	}

	return
}

func (r *repoHandler) GetCats(ctx context.Context, query string) (cats []entity.Cat, err error) {

	rows, err := r.catDB.Query(query)

	if err != nil {
		return
	}

	defer rows.Close()

	cat := entity.Cat{}

	for rows.Next() {

		err = rows.Scan(&cat.Id, &cat.Name, &cat.AgeInMonth, &cat.Race, &cat.Sex, &cat.Description, &cat.CreatedAt)

		if err != nil {
			return
		}

		cats = append(cats, cat)
	}

	return
}
