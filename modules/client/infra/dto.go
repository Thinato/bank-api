package client_repository

type Client struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Investor bool   `json:"investor"`
	Archived bool   `json:"archived"`
}