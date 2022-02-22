package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"net/http"
)

func main() {
	config.InitializerVariableEnvironment()
	fmt.Println(config.Port)
	fmt.Println("########## Welcome to API DevBooks ###########")
	fmt.Println("########## Server is running on port 5000 ###########")

	r := router.CreateRouter()
	err := http.ListenAndServe("localhost:5000", r)
	if err != nil {
		return
	}

}
