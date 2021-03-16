package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"net/http"
)

type object struct {
}

func newObject() *object {
	return &object{}
}

func main() {
	fx.New(
		fx.Provide(newObject),
		fx.Invoke(register),
	).Run()
}

func (o *object) doStuff() {
	fmt.Println("test")
}
func doStuff(obj *object) {
	obj.doStuff()
}

func register(lifecycle fx.Lifecycle) {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		},
	)
}
