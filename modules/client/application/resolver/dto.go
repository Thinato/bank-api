package client_resolver

import (
	client_usecase_login "bank/modules/client/domain/usecase/login"
	client_usecase_signup "bank/modules/client/domain/usecase/signup"
)

type ClientResolver struct {
	signup client_usecase_signup.SignupUsecase
	login client_usecase_login.LoginUsecase
}