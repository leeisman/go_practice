package main

import "log"

type Test interface {
	Test1()
	Test2()
}

type test struct {
	GOOD string
}

func (t *test) Test1() {
	log.Print("test Test1")
}
func (t *test) Test2() {
	log.Print("test Test2")
}

type test2 struct {
	test
}

func (t *test2) Test1() {
	t.test.Test1()
	log.Print("test2 Test1")
}
func (t *test2) Test2() {
	t.test.Test2()
	log.Print("test2 Test2")
}
func main() {
	testObj := &test2{test:test{GOOD: "test"}}
	testObj.Test1()
	testObj.Test2()
}
