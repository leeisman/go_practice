package main

import (
	"context"
	geoip2 "github.com/oschwald/geoip2-golang"
	"golang_practice/get_ip"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"time"
)

func main() {

	//確認直連
	ch := make(chan bool)

	// url check
	go func() {
		url := "https://www.youtube.com/watch?v=J88RKgxoaGw&ab_channel=%E6%BB%BE%E7%9F%B3%E5%94%B1%E7%89%87ROCKRECORDS"
		log.Print("checkDirectConnection check start Direct")
		ctx, _ := context.WithTimeout(context.Background(), 8*time.Second)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			log.Print("req err: ", err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp == nil {
			log.Print("err: ", err)
			log.Print("resp: ", resp)
			return
		}
		log.Print("resp status: ", resp.Status, resp.StatusCode == http.StatusOK)
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			log.Print("url check")
			ch <- true
		} else {
			ch <- false
		}
	}()

	// ip check
	go func() {
		ip := get_ip.GetIP()
		isCN, err := isCN(ip)
		log.Print("isCN: ", isCN)
		if err != nil {
			return
		}
		if !isCN {
			log.Print("ip check")
			ch <- true
		}
	}()

	isDirectLink := <-ch
	log.Print("isDirectLink: ", isDirectLink)
}

func isCN(ip string) (bool, error) {
	geoLite, err := filepath.Abs("./check_china/GeoLite2-Country.mmdb")
	db, err := geoip2.Open(geoLite)
	if err != nil {
		log.Print("open err: ", err)
		return false, err
	}
	netIP := net.ParseIP(ip)
	record, err := db.City(netIP)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	log.Println("debug geo country: ", record.Country.IsoCode)
	return record.Country.IsoCode == "CN", nil
}
