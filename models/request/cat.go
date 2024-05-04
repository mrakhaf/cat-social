package request

type UploadCat struct {
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	Race        string   `json:"race" validate:"required"`
	Sex         string   `json:"sex" validate:"required,oneof=male female"`
	AgeInMonth  int      `json:"ageInMonth" validate:"required,gte=1,lte=120082"`
	Description string   `json:"description" validate:"required,min=1,max=200"`
	ImageUrls   []string `json:"imageUrls" validate:"required,min=1,dive,url"`
}

type GetCatParam struct {
	Id         *string `query:"id" validate:"omitempty"`
	Limit      *int    `query:"limit" validate:"omitempty,gt=1"`
	Offset     *int    `query:"offset" validate:"omitempty,gt=1"`
	Race       *string `query:"race" validate:"omitempty"`
	Sex        *string `query:"sex" validate:"omitempty,oneof=male female"`
	HasMatch   *string `query:"hasMatched" validate:"omitempty,oneof=true false"`
	AgeInMonth *string `query:"ageInMonth" validate:"omitempty,oneof=>4 <4 4"`
	Owned      *string `query:"owned" validate:"omitempty,oneof=true false"`
	Search     *string `query:"search" validate:"omitempty"`
}

type MatchCat struct {
	MatchCatId string `json:"matchCatId" validate:"required"`
	UserCatId  string `json:"userCatId" validate:"required"`
	Message    string `json:"message" validate:"required,min=5,max=120"`
}

type ApproveRejectMatch struct {
	MatchId string `json:"matchId"`
}

type RejectMatch struct {
	MatchId string `json:"matchId"`
}
