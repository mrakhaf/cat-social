package dto

type SaveCatDto struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type MatchCatsDto struct {
	Id         string `json:"id"`
	MatchCatId string `json:"matchCatDetail"`
	UserCatId  string `json:"userCatDetail"`
	IssuedBy   string `json:"issuedBy"`
	Message    string `json:"message"`
	CreatedAt  string `json:"createdAt"`
}
