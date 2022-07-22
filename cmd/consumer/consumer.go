package main

import (
	"log"

	"ciandt.com/techPartnerReponseHandler/pkg/messaging"
)

func main() {
	msg, err := messaging.CreateNewMessaging()

	if err != nil {
		log.Fatalf("Error creating messaging %v", err)
		return
	}

	msg.Start()
}
