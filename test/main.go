package main

import "log"

func main() {
	file := "'/Users/frankieli/goProjects/src/gitlab.silkrode.com.tw/team_golang/mobile_lib/libs/test/universal_download/download/46e0833f-58dd-49c2-9c02-496f9ce79cee[15].mp4'\n"
	log.Print(file[1:len(file)-2])

	var test = []int{1, 23, 4, 5}
	log.Print(test)
	log.Print(test[:])
}
