package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

type Core struct {
	DoLock sync.Mutex
}

func (c *Core) Refresh() {
	c.DoLock.Lock()
	log.Print("refresh")
	for i := 1; i < 3; i++ {
		time.Sleep(time.Second * 1)
		log.Print("sleep ", strconv.Itoa(i),"\n")
	}
	c.DoLock.Unlock()
}

func (c *Core) Start() {
	log.Print("Wait Start")
	c.DoLock.Lock()
	log.Print("Start")
	c.DoLock.Unlock()
}

func (c *Core) Stop() {
	log.Print("Wait Stop")
	c.DoLock.Lock()
	log.Print("Stop")
	c.DoLock.Unlock()
}

func (c *Core) Start1() {
	log.Print("Wait Start1")
	c.DoLock.Lock()
	log.Print("Start1")
	c.DoLock.Unlock()
}

func (c *Core) Stop1() {
	log.Print("Wait Stop1")
	c.DoLock.Lock()
	log.Print("Stop1")
	c.DoLock.Unlock()
}

func main() {
	// 測試當一個lock key
	// 其他在爭搶的時候
	// 不一定先排隊的先搶到lock的key
	core := &Core{}
	for {
		go func() {
			core.Refresh()
		}()
		time.Sleep(time.Second * 1)
		go func() {
			core.Start()
		}()
		go func() {
			core.Stop()
		}()
		go func() {
			core.Start1()
		}()
		go func() {
			core.Stop1()
		}()
		time.Sleep(time.Second * 5)
		log.Print("_________________________________________________________________________________________")
		log.Print("")
		log.Print("")
		log.Print("")
	}
}
