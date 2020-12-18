package main

import (
	"fmt"
	"go.uber.org/fx"
)

type object struct {
}

func newObject() *object {
	return &object{}
}

func main() {
	fx.New(
		fx.Provide(newObject),
		fx.Invoke(doStuff),
	).Run()
}

func (o *object) doStuff() {
	fmt.Println("test")
}
func doStuff(obj *object) {
	obj.doStuff()
}
