package entity

import "time"

type User struct {
	Id       string
	Email    string
	Password string
	Name     string
}

type Cat struct {
	Id          string
	Name        string
	Race        string
	Sex         string
	AgeInMonth  int
	Description string
	ImageUrls   string
	HasMatched  bool
	CreatedAt   time.Time
}
