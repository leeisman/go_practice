package main

import "log"

func main() {
	var test = []int{1, 23, 4, 5}
	log.Print(test)
	log.Print(test[:])
}
