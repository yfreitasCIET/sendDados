package main

import (
	"fmt"
	"log"

	"ciandt.com/techPartnerReponseHandler/pkg/chat"
)

func main() {

	service, err := chat.NewTokenService()

	if err != nil {
		log.Fatal("Failed to create token service")
		return
	}
	token, err := service.GetToken()

	fmt.Println(*token)
	fmt.Println(service.GetToken())
	fmt.Println(service.GetToken())

}
