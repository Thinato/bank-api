package client

type ClientModule struct{}

func (module *ClientModule) Initialize() error {
	return nil
}

func CreateModule() IClientModule {
	return &ClientModule{}
}