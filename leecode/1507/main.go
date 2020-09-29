package main

import (
	"golang_practice/log_for_self/log"
	"strings"
)

func main() {
	date := ReformatDate("22nd Apr 2023")
	log.Print("date: ", date)
}

func ReformatDate(date string) string {
	daySplit := strings.Index(date, "th")
	if daySplit == -1 {
		daySplit = strings.Index(date, "nd")
	}
	if daySplit == -1 {
		daySplit = strings.Index(date, "rd")
	}
	if daySplit == -1 {
		daySplit = strings.Index(date, "st")
	}
	day := date[0:daySplit]
	if len(day) < 2 {
		day = "0" + day
	}
	year := date[len(date)-4:]
	month := date[daySplit+3 : daySplit+6]
	mapMonth := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
	value, ok := mapMonth[month]
	if ok {
		return year + "-" + value + "-" + day
	}
	return ""
}
