package client_usecase_login

import client_repository "bank/modules/client/infra"

type LoginUsecase struct {
	repository client_repository.IClientRepository
}

func (usecase *LoginUsecase) Run(input ClientLoginInputDto) (string, error) {
	
	
	return "", nil
}

func CreateLoginUsecase(
	repository client_repository.IClientRepository,
) ILoginUsecase {
	return &LoginUsecase{
		repository: repository,
	}
}