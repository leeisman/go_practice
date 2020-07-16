package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	do()
	//do()
}

func do() {
	URL := "https://content34.im.tv/jungle/recording/288265560523833784/watermark.mp4"
	ctx := context.Background()
	//tr := &http.Transport{    //解决x509: certificate signed by unknown authority
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	client := &http.Client{
		//Transport: tr,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	res, err := client.Do(req)
	if err != nil {
		log.Print("res err: ", err)
		return
	}
	log.Print("first res: ", res)
}
