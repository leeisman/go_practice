package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	pause int32
	wg sync.WaitGroup
}

type Downloader struct {
	cfg Config
}

func (d *Downloader) Pause()  {
	result := atomic.AddInt32(&d.cfg.pause, 1)
	if result <= 1 {
		d.cfg.wg.Add(1)
		fmt.Println("Pause !!!")
	}
}

func (d *Downloader) Resume()  {
	atomic.StoreInt32(&d.cfg.pause, 0)
	d.cfg.wg.Done()
	fmt.Println("Resume")
}


func  Download(cfg Config)  {
	for {
		cfg.wg.Wait()
		select {
		default:
			fmt.Println("Downloading ...")
			time.Sleep(3*time.Second)
		}
	}
}


func main() {
	d := Downloader{Config{
	}}
	go func() {
		for {
			d.Pause()
			time.Sleep(time.Second)
			d.Resume()
			time.Sleep(time.Second)
		}
	}()
	Download(d.cfg)
}


