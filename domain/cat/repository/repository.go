package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/mrakhaf/cat-social/domain/cat/interfaces"
	"github.com/mrakhaf/cat-social/models/dto"
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
	createdAt := utils.GenerateDatetimeIso8601()

	querySaveCat := fmt.Sprintf(`INSERT INTO cats (id, name, ageInMonths, race, sex, description, userId, createdAt) VALUES ('%s', '%s', %d, '%s', '%s', '%s', '%s', '%s')`, idCat, req.Name, req.AgeInMonth, req.Race, req.Sex, req.Description, userId, createdAt)

	querySaveImageCat := `INSERT INTO cat_images (id, idCat, imageUrl) VALUES ('%s', '%s', '%s')`

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

	for _, image := range req.ImageUrls {
		id := utils.GenerateUUID()
		_, err = tx.Exec(fmt.Sprintf(querySaveImageCat, id, idCat, image))

		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
		return
	}

	data = dto.SaveCatDto{
		Id:        idCat,
		CreatedAt: createdAt,
	}

	return
}
