package pool

import (
	"context"
	"fmt"
	"golang_practice/custom_log/log"
	"time"
)

var DefaultWorker = NewWorker(context.Background(), 2)

func example() {
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
	Cancel  context.CancelFunc
}

func NewWorker(ctx context.Context, size int) *Worker {
	ctx, cancel := context.WithCancel(ctx)
	w := &Worker{
		chanJob: make(chan *Job),
		size:    size,
		Cancel:  cancel,
	}
	go w.work(ctx)
	return w
}

func (w *Worker) DoAsync(job *Job) {
	w.chanJob <- job
}

func (w *Worker) Do(job *Job) {
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
	for i := 1; i <= w.size; i++ {
		go func(workerNo int) {
			for {
				select {
				case <-ctx.Done():
					log.Print("workerNo: ", workerNo, " done!")
					return
				case job := <-w.chanJob:
					log.Print("workerNo: ", workerNo, " doing job!")
					job.fn(ctx)
					log.Print("workerNo: ", workerNo, " job done!")
					close(job.finish)
				}
			}
		}(i)
	}
}
