package main

import "log"

type DirectHandler interface {
	InjectBillManager(BillManager)
	BillManager() BillManager
}
type BillManager interface {
	Start()
	Stop()
	SendSign(int)
	WaitReceiveSign()
}

type DirectI struct {
	billManger BillManager
}

func (d *DirectI) InjectBillManager(i BillManager) {
	d.billManger = i
}

func (d *DirectI) BillManager() BillManager {
	return d.billManger
}

type BillI struct {
	cost        int
	receiveSign chan int
}

func (b *BillI) Start() {
	log.Print("start: ", b.cost)
	b.cost += 20
}

func (b *BillI) Stop() {
	log.Print("stop ", b.cost)
	b.cost += 30
}

func (b *BillI) SendSign(sign int) {
	log.Print("before sign: ", b.receiveSign)
	b.receiveSign <- sign
}

func (b *BillI) WaitReceiveSign() {
	go func(b *BillI) {
		for {
			select {
			case ch := <-b.receiveSign:
				log.Print("receive sign: ", ch)
			}
		}
	}(b)
}

func main() {
	directHandler := &DirectI{}
	billManager := &BillI{
		receiveSign: make(chan int, 0),
	}
	billManager.WaitReceiveSign()
	billManager.SendSign(10)
	billManager.SendSign(20)
	directHandler.InjectBillManager(billManager)
	directHandler.BillManager().Start()
	directHandler.BillManager().Start()
	directHandler.BillManager().Start()
	directHandler.BillManager().Start()
}
