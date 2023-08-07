package client_repository

import (
	"bank/database"
	"bank/modules/client/domain/entities"
	client_usecase_signup "bank/modules/client/domain/usecase/signup"
)

type ClientRepository struct {
	conn *database.Database
}

// Deposit implements IClientRepository.
func (repo *ClientRepository) Deposit(token string, amount float32) error {
	panic("unimplemented")
}

// GetBalance implements IClientRepository.
func (repo *ClientRepository) GetBalance(token string) error {
	panic("unimplemented")
}

// GetTransactions implements IClientRepository.
func (repo *ClientRepository) GetTransactions(token string) error {
	panic("unimplemented")
}

// Invest implements IClientRepository.
func (repo *ClientRepository) Invest() error {
	panic("unimplemented")
}

// Login implements IClientRepository.
func (repo *ClientRepository) Login(username string, password string) (string, error) {
	panic("unimplemented")
}

// Signup implements IClientRepository.
func (repo *ClientRepository) Signup(username string, password string) error {
	usecase := client_usecase_signup.CreateSignupUsecase(repo)
	input := client_usecase_signup.ClientSignupInputDto{
		Username: username,
		Password: password,
	}
	err := usecase.Run(input)

	
	if err != nil {
		println("signup failed >:(")
		return err
	}
	return nil
}

func (repo *ClientRepository) SaveNewClient(entities.Client) error {
	_, err := repo.conn.
	return err
}

// Transfer implements IClientRepository.
func (repo *ClientRepository) Transfer(token string, recipient Client, amount float32) error {
	panic("unimplemented")
}

// Withdraw implements IClientRepository.
func (repo *ClientRepository) Withdraw(token string, amount float32) error {
	panic("unimplemented")
}

func CreateClientRepository() IClientRepository {
	return &ClientRepository{}
}
