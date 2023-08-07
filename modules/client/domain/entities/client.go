package entities

import "time"

type Client struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Investor  bool      `json:"investor"`
	Archived  bool      `json:"archived"`
	Balance   float32   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}