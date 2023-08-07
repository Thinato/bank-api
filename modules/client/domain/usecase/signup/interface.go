package client_usecase_signup

type ISignupUsecase interface {
	Run(input ClientSignupInputDto) error
}