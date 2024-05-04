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

		err = rows.Scan(&cat.Id, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.HasMatched, &cat.ImageUrls, &cat.CreatedAt)

		if err != nil {
			return
		}

		cats = append(cats, cat)
	}

	return
}

func (r *repoHandler) GetCatUser(ctx context.Context, userId, catId string) (cats entity.Cat, err error) {

	query := fmt.Sprintf(`SELECT id FROM cats WHERE userId = '%s' AND id = '%s'`, userId, catId)

	fmt.Println(query)

	row := r.catDB.QueryRow(query)

	err = row.Scan(&cats.Id)

	if err != nil {
		return
	}

	return

}

func (r *repoHandler) UpdateCat(ctx context.Context, catId, userId string, req request.UploadCat) (err error) {
	imageUrls := utils.SliceToString(req.ImageUrls)

	query := fmt.Sprintf(`UPDATE cats SET name = '%s', ageInMonths = %d, race = '%s', sex = '%s', description = '%s', imageUrls = '%s' WHERE id = '%s' AND userId = '%s'`, req.Name, req.AgeInMonth, req.Race, req.Sex, req.Description, imageUrls, catId, userId)

	_, err = r.catDB.Exec(query)

	if err != nil {
		return
	}

	return
}

func (r *repoHandler) DeleteCat(ctx context.Context, catId string) (err error) {

	query := fmt.Sprintf(`DELETE FROM cats WHERE id = '%s'`, catId)

	_, err = r.catDB.Exec(query)

	if err != nil {
		return
	}

	return
}

func (r *repoHandler) GetCatUserHasNotMatch(ctx context.Context, userId, catId string) (sex string, err error) {

	var cat entity.Cat

	query := fmt.Sprintf(`SELECT id, sex FROM cats WHERE userId = '%s' AND id = '%s' AND hasMatched IS NOT TRUE`, userId, catId)
	fmt.Println(query)
	row := r.catDB.QueryRow(query)

	err = row.Scan(&cat.Id, &cat.Sex)

	if err != nil {
		return
	}

	sex = cat.Sex

	return

}

func (r *repoHandler) GetCatMatchHasNotMatch(ctx context.Context, userId, catId string) (sex string, err error) {

	var cat entity.Cat

	query := fmt.Sprintf(`SELECT id, sex FROM cats WHERE userId != '%s' AND id = '%s' AND hasMatched IS NOT TRUE`, userId, catId)
	fmt.Println(query)
	row := r.catDB.QueryRow(query)

	err = row.Scan(&cat.Id, &cat.Sex)

	if err != nil {
		return
	}

	sex = cat.Sex

	return

}

func (r *repoHandler) SaveMatchCat(ctx context.Context, userId string, req request.MatchCat) (matchId string, err error) {

	matchId = utils.GenerateUUID()

	query := fmt.Sprintf(`INSERT INTO match_cats (matchid, matchcatid, usercatid, status, issuedby, message, createdat) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s')`, matchId, req.MatchCatId, req.UserCatId, "pending", userId, req.Message, time.Now().Format(time.RFC3339))

	_, err = r.catDB.Exec(query)

	if err != nil {
		return
	}

	return

}

func (r *repoHandler) GetCatByID(ctx context.Context, catId string) (cat entity.Cat, err error) {

	query := fmt.Sprintf(`SELECT id, name, race, sex, ageInMonths, description, imageUrls, hasMatched, createdAt FROM cats WHERE id = '%s'`, catId)

	row := r.catDB.QueryRow(query)

	err = row.Scan(&cat.Id, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.ImageUrls, &cat.HasMatched, &cat.CreatedAt)

	if err != nil {
		return
	}

	return
}

func (r *repoHandler) GetMatchByIdAndUserIdApprover(ctx context.Context, matchId string, userId string) (matchCatId, userCatId string, err error) {

	query := fmt.Sprintf(`SELECT matchid, matchcatid, usercatid FROM match_cats 
		INNER JOIN cats on match_cats.matchcatid = cats.id 
		WHERE matchid = '%s' 
		AND userid = '%s' AND status = 'pending'`, matchId, userId)

	fmt.Println(query)

	row := r.catDB.QueryRow(query)

	err = row.Scan(&matchId, &matchCatId, &userCatId)

	if err != nil {
		return
	}

	return

}

func (r *repoHandler) GetMatchStatusPending(ctx context.Context, matchId string) (err error) {

	query := fmt.Sprintf(`SELECT matchid FROM match_cats WHERE status = 'pending' AND matchid = '%s'`, matchId)

	fmt.Println(query)

	row := r.catDB.QueryRow(query)

	err = row.Scan(&matchId)

	if err != nil && err == sql.ErrNoRows {
		err = nil
		return
	}

	return

}

func (r *repoHandler) UpdatedMatchStatus(ctx context.Context, matchId, status, matchCatId, userCatId string) (err error) {

	queryUpdateStatus := fmt.Sprintf(`UPDATE match_cats SET status = '%s' WHERE matchid = '%s'`, status, matchId)
	queryUpdateUserCatHasMatch := fmt.Sprintf(`UPDATE cats SET hasMatched = true WHERE id = '%s'`, userCatId)
	queryUpdateMatchCatHasMatch := fmt.Sprintf(`UPDATE cats SET hasMatched = true WHERE id = '%s'`, matchCatId)

	tx, err := r.catDB.Begin()

	if err != nil {
		return
	}

	_, err = tx.Exec(queryUpdateStatus)

	if err != nil {
		tx.Rollback()
		return
	}

	_, err = tx.Exec(queryUpdateUserCatHasMatch)

	if err != nil {
		tx.Rollback()
		return
	}

	_, err = tx.Exec(queryUpdateMatchCatHasMatch)

	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	if err != nil {
		return
	}

	return

}
