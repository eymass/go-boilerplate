package main

import (
	"fmt"
	"github.com/heroku/go-getting-started/auth"
	"github.com/heroku/go-getting-started/service"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	secretKey, err := auth.GenerateSecretKey(32) // 32 bytes = 256 bits
	if err != nil {
		panic(err)
	}
	fmt.Println("Secret Key:", secretKey)

	serviceInstance := service.InitService()
	serviceInstance.Start()
}
