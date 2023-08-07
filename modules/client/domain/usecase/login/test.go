package client_usecase_login

type MockClientRepository struct {
}

func (repository *MockClientRepository) Login() (string, error) {
	// get token in database
	return "", nil
}