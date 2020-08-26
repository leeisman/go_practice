package main

import (
	"context"
	"golang_practice/custom_log/log"
	"golang_practice/worker/pool"
	"sync"
)

func main() {
	p1 := "p1"
	p2 := 2
	wg := &sync.WaitGroup{}
	fn := func(p1 string, p2 int, wg *sync.WaitGroup) func(ctx context.Context) {
		return func(ctx context.Context) {
			defer wg.Done()
			log.Print("p1: ", p1)
			log.Print("p2: ", p2)
		}
	}

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		pool.DefaultWorker.DoAsync(pool.NewJob(fn(p1, p2, wg)))
	}

	wg.Wait()
}
