package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"net/http"
)

//func init() {
//	key := make([]byte, 64)
//	if _, err := rand.Read(key); err != nil {
//		log.Fatal(err)
//	}
//	base64 := base642.StdEncoding.EncodeToString(key)
//	fmt.Println(base64)
//}

func main() {
	config.InitializerVariableEnvironment()
	fmt.Println(config.Port)
	fmt.Println("########## Welcome to API DevBooks ###########")
	fmt.Println("########## Server is running on port 5000 ###########")

	r := router.CreateRouterUser()
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)
	if err != nil {
		return
	}

}
