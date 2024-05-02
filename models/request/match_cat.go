package request

type MatchCat struct {
	MatchCatId string `json:"matchCatId" validate:"required"`
	UserCatId  string `json:"userCatId" validate:"required"`
}
