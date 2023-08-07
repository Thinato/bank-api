package main

import (
	"bank/database"
	"bank/modules/client"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	env := LoadEnv()
	router := gin.Default()

	db, err := database.NewDatabase()
    if err != nil {
        panic(err)
    }
    defer db.Close()

    clientModule := client.CreateModule(*db)
	err = clientModule.Initialize()
	if err != nil {
		panic("could not initialize module")
	}


	fmt.Println("Starting bank. . .")

	err = router.Run("127.0.0.1:" + env.Port)
	defer db.Close()
	if err != nil {
		return
	}
}
