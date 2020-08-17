package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	w := NewWorker(ctx, 2)
	i := 0
	w.DoAsync(NewJob(func(ctx context.Context) {
		i++
		fmt.Println("!!!!! time.Sleep(3 * time.Second)")
		time.Sleep(3 * time.Second)
	}))
	w.DoAsync(NewJob(func(ctx context.Context) {
		i++
		fmt.Println("!!!!! time.Sleep(3 * time.Second)")
		time.Sleep(3 * time.Second)
	}))
	w.DoAsync(NewJob(func(ctx context.Context) {
		i++
		fmt.Println("!!!!! time.Sleep(3 * time.Second)")
		time.Sleep(3 * time.Second)
	}))
	w.DoAsync(NewJob(func(ctx context.Context) {
		i++
		fmt.Println("!!!!! time.Sleep(3 * time.Second)")
		time.Sleep(3 * time.Second)
	}))
	fmt.Println(i)
	time.Sleep(5 * time.Second)
}

type Worker struct {
	size    int
	chanJob chan *Job
}

func NewWorker(ctx context.Context, size int) *Worker {
	w := &Worker{
		chanJob: make(chan *Job),
		size:    size,
	}
	go w.work(ctx)
	return w
}

func (w *Worker) DoAsync(job *Job) {
	w.chanJob <- job
}

func (w *Worker) DoSync(job *Job) {
	w.chanJob <- job
	<-job.finish
}

type Job struct {
	fn     func(ctx context.Context)
	finish chan struct{}
}

func NewJob(fn func(ctx context.Context)) *Job {
	return &Job{
		fn:     fn,
		finish: make(chan struct{}),
	}
}

func (w *Worker) work(ctx context.Context) {
	for i := 0; i < w.size; i++ {
		go func() {
			for job := range w.chanJob {
				select {
				case <-ctx.Done():
				default:
					job.fn(ctx)
				}
				close(job.finish)
			}
		}()
	}
}
