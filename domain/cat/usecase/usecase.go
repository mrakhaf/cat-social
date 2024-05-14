package usecase

import (
	"context"
	"fmt"
	"strings"

	Auth "github.com/mrakhaf/cat-social/domain/auth/interfaces"
	Cat "github.com/mrakhaf/cat-social/domain/cat/interfaces"
	"github.com/mrakhaf/cat-social/models/request"
	"github.com/mrakhaf/cat-social/models/response"
)

type usecase struct {
	catRepo  Cat.Repository
	authRepo Auth.Repository
}

func NewUsecase(catRepo Cat.Repository, authRepo Auth.Repository) Cat.Usecase {
	return &usecase{
		catRepo:  catRepo,
		authRepo: authRepo,
	}
}

func (u *usecase) UploadCat(ctx context.Context, req request.UploadCat, userId string) (data interface{}, err error) {

	data, err = u.catRepo.SaveCat(ctx, req, userId)

	if err != nil {
		err = fmt.Errorf("failed to save cat: %s", err)
		return
	}

	return

}

func (u *usecase) GetCat(ctx context.Context, userId string, req request.GetCatParam) (data interface{}, err error) {

	query := "SELECT id, name, race, sex, ageinmonths, description, hasmatched, imageurls, createdat FROM cats"

	var firstFilterParam bool

	if req.Id != nil {
		query = fmt.Sprintf(` %s WHERE id = '%s'`, query, *req.Id)
		firstFilterParam = true
	}

	if req.Race != nil {
		if firstFilterParam {
			query = fmt.Sprintf(` %s AND race = '%s'`, query, *req.Race)
		} else {
			query = fmt.Sprintf(` %s WHERE race = '%s'`, query, *req.Race)
			firstFilterParam = true
		}

	}

	if req.Sex != nil {
		if firstFilterParam {
			query = fmt.Sprintf(` %s AND sex = '%s'`, query, *req.Sex)
		} else {
			query = fmt.Sprintf(` %s WHERE sex = '%s'`, query, *req.Sex)
			firstFilterParam = true
		}
	}

	if req.HasMatch != nil {
		if firstFilterParam {
			if *req.HasMatch == "true" {
				query = fmt.Sprintf(` %s AND hasMatched IS TRUE`, query)
			} else if *req.HasMatch == "false" {
				query = fmt.Sprintf(` %s AND hasMatched IS NOT TRUE`, query)
			}
		} else {
			if *req.HasMatch == "true" {
				query = fmt.Sprintf(` %s WHERE hasMatched IS TRUE`, query)
				firstFilterParam = true
			} else if *req.HasMatch == "false" {
				query = fmt.Sprintf(` %s WHERE hasMatched IS NOT TRUE`, query)
				firstFilterParam = true
			}
		}
	}

	if req.AgeInMonth != nil {
		if firstFilterParam {

			if *req.AgeInMonth == "<4" {
				query = fmt.Sprintf(` %s AND ageinmonths < 4`, query)
			} else if *req.AgeInMonth == "4" {
				query = fmt.Sprintf(` %s AND ageinmonths = 4`, query)
			} else if *req.AgeInMonth == ">4" {
				query = fmt.Sprintf(` %s AND ageinmonths > 4`, query)
			}
		} else {
			if *req.AgeInMonth == "<4" {
				query = fmt.Sprintf(` %s WHERE ageinmonths < 4`, query)
				firstFilterParam = true
			} else if *req.AgeInMonth == "4" {
				query = fmt.Sprintf(` %s WHERE ageinmonths = 4`, query)
				firstFilterParam = true
			} else if *req.AgeInMonth == ">4" {
				query = fmt.Sprintf(` %s WHERE ageinmonths > 4`, query)
				firstFilterParam = true
			}
		}
	}

	if req.Owned != nil {
		if firstFilterParam {

			if *req.Owned == "true" {
				query = fmt.Sprintf(` %s AND userId = '%s'`, query, userId)
			} else if *req.Owned == "false" {
				query = fmt.Sprintf(` %s AND userId != '%s'`, query, userId)
			}

		} else {
			if *req.Owned == "true" {
				query = fmt.Sprintf(` %s WHERE userId = '%s'`, query, userId)
				firstFilterParam = true
			} else if *req.Owned == "false" {
				query = fmt.Sprintf(` %s WHERE userId != '%s'`, query, userId)
				firstFilterParam = true
			}
		}
	}

	if req.Search != nil {
		if firstFilterParam {
			query = fmt.Sprintf(` %s AND name LIKE '%%%s%%'`, query, *req.Search)
		} else {
			query = fmt.Sprintf(` %s WHERE name LIKE '%%%s%%'`, query, *req.Search)
		}
	}

	if req.Limit != nil {
		query = fmt.Sprintf(` %s LIMIT %d`, query, *req.Limit)
	}

	if req.Offset != nil {
		query = fmt.Sprintf(` %s OFFSET %d`, query, *req.Offset)
	}

	fmt.Println(query)

	cats, err := u.catRepo.GetCats(ctx, query)

	if err != nil {
		err = fmt.Errorf("failed to get cats: %s", err)
		return
	}

	catsResponse := []response.GetCats{}

	for _, cat := range cats {
		imageUrls := strings.Split(cat.ImageUrls, ",")
		createdAt := cat.CreatedAt.Format("2006-01-02")

		catsResponse = append(catsResponse, response.GetCats{
			Id:          cat.Id,
			Name:        cat.Name,
			Race:        cat.Race,
			Sex:         cat.Sex,
			AgeInMonths: cat.AgeInMonth,
			Description: cat.Description,
			HasMatched:  cat.HasMatched,
			ImageUrls:   imageUrls,
			CreatedAt:   createdAt,
		})
	}

	data = catsResponse

	return

}

func (u *usecase) ValidationCatUser(ctx context.Context, catId, userId string) (err error) {
	_, err = u.catRepo.GetCatUser(ctx, userId, catId)

	if err != nil {
		err = fmt.Errorf("failed to get cat: %s", err)
		return
	}

	return
}

func (u *usecase) UpdateCat(ctx context.Context, catId, userId string, req request.UploadCat) (err error) {

	err = u.catRepo.UpdateCat(ctx, catId, userId, req)

	if err != nil {
		err = fmt.Errorf("failed to update cat: %s", err)
		return
	}

	return
}

func (u *usecase) DeleteCat(ctx context.Context, catId string) (err error) {

	err = u.catRepo.DeleteCat(ctx, catId)

	if err != nil {
		err = fmt.Errorf("failed to delete cat: %s", err)
		return
	}

	return

}

func (u *usecase) ValidateMatchCat(ctx context.Context, userId string, req request.MatchCat) (err error) {

	//get cat
	_, err = u.catRepo.GetCatByID(ctx, req.MatchCatId)

	if err != nil {
		err = fmt.Errorf("404")
		return
	}

	//get cat user
	_, err = u.catRepo.GetCatUser(ctx, userId, req.UserCatId)

	if err != nil {
		fmt.Println(err.Error())
		err = fmt.Errorf("404")
		return
	}

	//get cat user not matched
	sexUserCat, err := u.catRepo.GetCatUserHasNotMatch(ctx, userId, req.UserCatId)

	if err != nil {
		err = fmt.Errorf("400")
		return
	}

	//get cat match no matched
	sexMatchCat, err := u.catRepo.GetCatMatchHasNotMatch(ctx, userId, req.MatchCatId)

	if err != nil {
		err = fmt.Errorf("400")
		return
	}

	fmt.Println(sexUserCat)
	fmt.Println(sexMatchCat)
	//validate sex
	if sexUserCat == sexMatchCat {
		err = fmt.Errorf("400")
		return
	}

	return

}

func (u *usecase) UploadMatch(ctx context.Context, req request.MatchCat, userId string) (err error) {

	_, err = u.catRepo.SaveMatchCat(ctx, userId, req)

	if err != nil {
		err = fmt.Errorf("failed to upload match: %s", err)
		return
	}

	return

}

func (u *usecase) ApproveMatch(ctx context.Context, req request.ApproveRejectMatch, matchCatId, userCatId string) (err error) {

	err = u.catRepo.UpdatedMatchStatus(ctx, req.MatchId, "approved", matchCatId, userCatId)

	if err != nil {
		err = fmt.Errorf("failed to approve match: %s", err)
		return
	}

	return

}

func (u *usecase) RejectMatch(ctx context.Context, req request.RejectMatch, matchCatId, userCatId string) (err error) {

	err = u.catRepo.UpdatedMatchStatus(ctx, req.MatchId, "rejected", matchCatId, userCatId)

	if err != nil {
		err = fmt.Errorf("failed to reject match: %s", err)
		return
	}

	return
}

func (u *usecase) DeleteMatch(ctx context.Context, id string, userId string) (err error) {

	matchCatId := u.catRepo.GetMatchById(ctx, id)

	if matchCatId != nil {
		err = fmt.Errorf("404")
		return
	}

	status, err := u.catRepo.GetMatchStatus(ctx, id)

	if err != nil {
		err = fmt.Errorf("failed to get match status: %s", err)
		return
	}

	if status == "approved" || status == "rejected" {
		err = fmt.Errorf("400")
		return
	}

	err = u.catRepo.DeleteMatch(ctx, id, userId)

	if err != nil {
		err = fmt.Errorf("failed to delete match: %s", err)
		return
	}

	return
}

func (u *usecase) GetMatchs(ctx context.Context, userId string) (data interface{}, err error) {

	matchs, err := u.catRepo.GetAllMatchByUserId(ctx, userId)

	if err != nil {
		err = fmt.Errorf("failed to get matchs: %s", err)
		return
	}

	issuedby, err := u.authRepo.GetDataAccount(matchs.IssuedBy)

	if err != nil {
		err = fmt.Errorf("failed to get data account: %s", err)
		return
	}

	matchCatDetail, err := u.catRepo.GetCatByID(ctx, matchs.MatchCatId)

	if err != nil {
		err = fmt.Errorf("failed to get cat detail: %s", err)
		return
	}

	userCatDetail, err := u.catRepo.GetCatByID(ctx, matchs.UserCatId)
	if err != nil {
		err = fmt.Errorf("failed to get cat detail: %s", err)
		return
	}

	data = response.GetMatchs{
		Id: matchs.Id,
		IssuedBy: response.IssuedBy{
			Name:      issuedby.Name,
			Email:     issuedby.Email,
			CreatedAt: issuedby.CreatedAt.Format("2006-01-02"),
		},
		MatchCatDetail: response.CatDetail{
			Id:          matchCatDetail.Id,
			Name:        matchCatDetail.Name,
			Race:        matchCatDetail.Race,
			Sex:         matchCatDetail.Sex,
			Description: matchCatDetail.Description,
			AgeInMonth:  matchCatDetail.AgeInMonth,
			ImageUrls:   strings.Split(matchCatDetail.ImageUrls, ","),
			HasMatched:  matchCatDetail.HasMatched,
			CreatedAt:   matchCatDetail.CreatedAt.Format("2006-01-02"),
		},
		UserCatDetail: response.CatDetail{
			Id:          userCatDetail.Id,
			Name:        userCatDetail.Name,
			Race:        userCatDetail.Race,
			Sex:         userCatDetail.Sex,
			Description: userCatDetail.Description,
			AgeInMonth:  userCatDetail.AgeInMonth,
			ImageUrls:   strings.Split(userCatDetail.ImageUrls, ","),
			HasMatched:  userCatDetail.HasMatched,
			CreatedAt:   userCatDetail.CreatedAt.Format("2006-01-02"),
		},
		Message:   matchs.Message,
		CreatedAt: matchs.CreatedAt,
	}

	return data, nil
}
