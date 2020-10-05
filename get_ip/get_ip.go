package get_ip

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetIP() string {
	url := "https://api.ipify.org?format=text"
	fmt.Printf("Getting IP address from  ipify\n")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("My IP is:%s\n", ip)
	return string(ip)
}
