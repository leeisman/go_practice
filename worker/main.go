package main

import (
	"context"
	worker "golang_practice/worker/pool"
	"log"
	"sync"
	"time"
)

func main() {
	go func() {
		log.Print("!!!!!! download task")
		worker.SetSize(2)
		wg := &sync.WaitGroup{}

		for i := 0; i < 10; i++ {
			wg.Add(1)
			fn := func(i int) func(ctx context.Context) {
				return func(ctx context.Context) {
					time.Sleep(time.Second * 5)
					log.Print("download ts ", i)
					wg.Done()
				}
			}
			job := worker.NewJob(fn(i))
			worker.DefaultWorker.DoAsync(job)
			log.Print("add download job ", i)
		}
		wg.Wait()
	}()
	time.Sleep(time.Second * 5)
	log.Print("!!!!!! get info task")
	go func() {
		wg := &sync.WaitGroup{}
		for i := 0; i < 2; i++ {
			wg.Add(1)
			fn := func(i int) func(ctx context.Context) {
				return func(ctx context.Context) {
					log.Print("get info ", i)
					wg.Done()
				}
			}
			job := worker.NewJob(fn(i))
			worker.DefaultWorker.DoAsync(job)
		}
		wg.Wait()
	}()

	time.Sleep(time.Second * 1000)
}
