package client

import "bank/database"

type ClientModule struct{}

func (module *ClientModule) Initialize() error {
	return nil
}

func CreateModule(db database.Database) IClientModule {
	return &ClientModule{}
}