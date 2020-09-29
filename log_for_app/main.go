package main

import "golang_practice/log_for_app/log"

func main() {
	log.ErrorMode()
	log.DebugMode()
	log.InfoMode()
	log.Print("test", "test2", "test3")
	log.DebugMsg("test", "test2")
	log.ErrorMsg("test", "test2")
	log.Print("test", "test2", "test4")
}