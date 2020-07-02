package main

import (
	"log"
	"time"
)

func main() {
	isLock := true
	ch1 := make(chan int, 0)
	ch2 := make(chan int, 0)
	ch3 := make(chan int, 0)

	if isLock {
		go func() {
			// 沒default 就會lock
			for {
				log.Print("have lock loop start")
				select {
				case ch := <-ch1:
					log.Print("ch signal: ", ch)
					log.Print("ch: ", ch)
				case ch := <-ch2:
					log.Print("ch signal: ", ch)
					log.Print("ch: ", ch)
				case ch := <-ch3:
					log.Print("ch signal: ", ch)
					log.Print("ch: ", ch)
				}
				time.Sleep(time.Second * 1)
				log.Print("have lock loop end")
			}
		}()
	} else {
		go func() {
			// 有default 就不lock
			for {
				log.Print("no lock loop start")
				select {
				case ch := <-ch1:
					log.Print("ch signal: ", ch)
					log.Print("ch: ", ch)
				case ch := <-ch2:
					log.Print("ch signal: ", ch)
					log.Print("ch: ", ch)
				case ch := <-ch3:
					log.Print("ch signal: ", ch)
					log.Print("ch: ", ch)
				default:
					log.Print("default")
				}
				time.Sleep(time.Second * 1)
				log.Print("no lock loop end")
			}
		}()
	}

	go func() {
		for {
			log.Print("signal wait 5s")
			time.Sleep(time.Second * 5)
			go func() {
				log.Print("set signal 1 to ch1")
				ch1 <- 1
			}()
			go func() {
				log.Print("set signal 2 to ch2")
				ch2 <- 2
			}()
			go func() {
				log.Print("set signal 3 to ch3")
				ch3 <- 3
			}()
		}
	}()

	time.Sleep(time.Second * 500000)

}
