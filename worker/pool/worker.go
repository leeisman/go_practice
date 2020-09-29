package worker

import (
	"context"
	"github.com/rs/zerolog"
	"log"
)

var DefaultWorker = NewWorker(context.Background(), 100)
var logger = zerolog.New(log.Writer()).With().Timestamp().Str("tag", "go-log").Logger()

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
					//logger.Print("workerNo: ", workerNo, " done!")
					return
				case job := <-w.chanJob:
					//logger.Print("workerNo: ", workerNo, " doing job!")
					job.fn(ctx)
					//logger.Print("workerNo: ", workerNo, " job done!")
					close(job.finish)
				}
			}
		}(i)
	}
}

func SetLogger(ll zerolog.Logger) {
	logger = ll
}

func SetSize(size int) {
	DefaultWorker.Cancel()
	logger.Print("worker cancel all job")
	DefaultWorker = NewWorker(context.Background(), size)
}
