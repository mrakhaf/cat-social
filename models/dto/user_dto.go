package dto

type AuthResponse struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
