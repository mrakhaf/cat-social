package response

type GetCats struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	AgeInMonths int      `json:"ageInMonth"`
	Description string   `json:"description"`
	ImageUrls   []string `json:"imageUrls"`
	HasMatched  bool     `json:"hasMatched"`
	CreatedAt   string   `json:"createdAt"`
}

type GetMatchs struct {
	Id             string    `json:"id"`
	IssuedBy       IssuedBy  `json:"issuedBy"`
	MatchCatDetail CatDetail `json:"matchCatDetail"`
	UserCatDetail  CatDetail `json:"userCatDetail"`
	Message        string    `json:"message"`
	CreatedAt      string    `json:"createdAt"`
}

type IssuedBy struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

type CatDetail struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	Description string   `json:"description"`
	AgeInMonth  int      `json:"ageInMonth"`
	ImageUrls   []string `json:"imageUrls"`
	HasMatched  bool     `json:"hasMatched"`
	CreatedAt   string   `json:"createdAt"`
}
