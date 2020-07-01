package main

import (
	"context"
	"log"
	"time"
)

func main() {

	// ctx帶入開關chan
	pauseCh := make(chan int, 0)
	startCh := make(chan int, 0)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "pause", pauseCh)
	ctx = context.WithValue(ctx, "start", startCh)

	// routine 開關狀態任務
	go func() {
		status := "start"
		for {
			time.Sleep(time.Millisecond * 500)
			ctx1 := ctx.Value("pause")
			ch1, ok1 := ctx1.(chan int)
			ctx2 := ctx.Value("start")
			ch2, ok2 := ctx2.(chan int)
			if ok1 == false {
				log.Print("ctx not exist pause channel")
			}
			if ok2 == false {
				log.Print("ctx not exist pause channel")
			}
			select {
			case <-ch1:
				log.Print("pause")
				status = "pause"
			case <-ch2:
				log.Print("start")
				status = "start"
			default:
				log.Print("now status ", status)
			}
		}
	}()

	// 切換成 pause event
	go func(duration time.Duration) {
		for {
			time.Sleep(time.Millisecond * 1500)
			ctxPause := ctx.Value("pause")
			pauseCh, ok := ctxPause.(chan int)
			if ok == false {
				log.Print("ctx not exist pause channel")
			}
			log.Print("1 to pause")
			pauseCh <- 1
			time.Sleep(time.Second * duration)
		}
	}(9)
	// 切換成 start event
	go func(duration time.Duration) {
		for {
			time.Sleep(time.Millisecond * 3500)
			ctxStart := ctx.Value("start")
			startCh, ok := ctxStart.(chan int)
			if ok == false {
				log.Print("ctx not exist pause channel")
			}
			log.Print("1 to start")
			startCh <- 1
			time.Sleep(time.Second * duration)
		}
	}(5)

	time.Sleep(time.Second * 100000)
}
