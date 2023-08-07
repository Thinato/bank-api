package client_usecase_signup

type ClientSignupInputDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}