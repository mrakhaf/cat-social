package dto

type AuthResponse struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccessToken string `json:"accessToken"`
}
