package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"golang_practice/easy_proxy"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
)

func main() {
	netIp := net.ParseIP(easy_proxy.IP)
	if netIp == nil {
		fmt.Print("ip invalid")
		return
	}
	httpProxy := fmt.Sprintf("http://%s%s", netIp, easy_proxy.Port)
	proxy, _ := url.Parse(httpProxy)
	http.DefaultClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}

	req, err := http.NewRequest("GET", "https://www.flysnow.org/2016/12/24/golang-http-proxy.html", nil)
	if err != nil {
		log.Error().Str("error", err.Error()).Msg("NewRequest")
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Str("error", err.Error()).Msg("client do")
		return
	}
	if resp == nil {
		log.Error().Msg("resp nil")
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Str("error", err.Error()).Msg("ioutil readAll")
	}
	log.Debug().Str("body", string(body)).Msg("response ok!!!!")
}
