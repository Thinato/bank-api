package client_usecase_login

type ILoginUsecase interface {
	Run(input ClientLoginInputDto) (string, error)
}