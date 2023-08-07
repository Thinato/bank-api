package client_usecase_login

type ClientLoginInputDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}