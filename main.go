package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	env := LoadEnv()
	router := gin.Default()

	fmt.Println("Starting bank. . .")

	err := router.Run("127.0.0.1:" + env.Port)

	if err != nil {
		return
	}
}
