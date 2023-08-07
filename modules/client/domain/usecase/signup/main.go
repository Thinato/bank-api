package client_usecase_signup

import (
	"bank/modules/client/domain/entities"
	client_repository "bank/modules/client/infra"
)

type SignupUsecase struct {
	repository client_repository.IClientRepository
}

func (usecase *SignupUsecase) Run(input ClientSignupInputDto) (error) {

	client := entities.Client{
        Username: input.Username,
        Password: input.Password,
    }
	return nil
}


func CreateSignupUsecase(
	repository client_repository.IClientRepository,
) ISignupUsecase {
	return &SignupUsecase {
		repository: repository,
	}
}