package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"ciandt.com/techPartnerReponseHandler/pkg/chat"
)

func main() {

	url := "https://chat.googleapis.com/v1/spaces/hQUbP4AAAAE/messages?threadKey=blablabla"
	body, _ := json.Marshal(map[string]string{
		"text": "VITORIA consgeue",
	})

	text, err := sendRequest(&http.Request{}, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Print(text.Status)
}

func sendRequest(r *http.Request, metodo, url string, body io.Reader) (*http.Response, error) {

	request, err := http.NewRequest(metodo, url, body)
	if err != nil {
		return nil, err
	}

	msg, err := chat.NewTokenService()
	if err != nil {
		log.Fatalf("%v", err)
	}

	token, err := msg.GetToken()
	if err != nil {
		log.Fatalf("%v", err)
	}

	request.Header.Add("Authorization", "Bearer "+string(*token))

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
