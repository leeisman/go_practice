package main

import (
	"fmt"
	"time"
)

type Config struct {
	Pause chan struct{}
	Resume chan struct{}
}

type Downloader struct {
	cfg Config
}

func (d *Downloader) Pause()  {
	d.cfg.Pause <- struct{}{}
}

func (d *Downloader) Resume()  {
	d.cfg.Resume <- struct{}{}
}


func  Download(cfg Config)  {
	for {
		select {
		case <-cfg.Pause:
			fmt.Println("Pause")
			select {
			case <-cfg.Resume:
				fmt.Println("Resume !!!!!")
			}
		default:
			fmt.Println("Downloading ...")
			time.Sleep(3*time.Second)
		}
	}
}


func main() {
	d := Downloader{Config{
		Pause: make(chan struct{}),
		Resume: make(chan struct{}),
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


