package model

type LoginRequestModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseModel struct {
	Token string `json:"token"`
}
