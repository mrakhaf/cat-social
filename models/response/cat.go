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
