package main

import "golang_practice/log_for_self/log"

func main() {
	log.ErrorMode()
	log.DebugMode()
	log.Print("test", "test2","test3")
	log.Debug("test", "test2")
	log.Error("test", "test2")
	log.Print("test", "test2","test4")

}
