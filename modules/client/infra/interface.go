package client_repository

import (
	"bank/modules/client/domain/entities"
)

type IClientRepository interface {
	Signup(username string, password string) error
	Login(username string, password string) (string, error)
	Deposit(token string, amount float32) error
	Withdraw(token string, amount float32) error
	Transfer(token string, recipient Client, amount float32) error
	Invest() error
	GetBalance(token string) error
	GetTransactions(token string) error
	SaveNewClient(client entities.Client)
}