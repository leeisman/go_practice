package main

import (
	"fmt"
	"log"
)

func main() {
	var num string
	var tail string
	name := "CONNECT www.flysnow.org:443 HTTP/1.1 Host: www.flysnow.org:443 User-Agent: Go-http-client/1.1"
	n, _ := fmt.Sscanf(name, "%s%S", &num, &tail)
	if n == 1 {
	}

	log.Print("num: ", num, " tail: ", tail)
}
