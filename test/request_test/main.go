package main

import (
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://yedgsb.com/ping", nil)
	if err != nil {
		log.Print("err: ", err)
		return
	}

	client := &http.Client{
		Transport: &http.Transport{
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Print("err: ", err)
		return
	}

	log.Print("resp: ", resp)
}
