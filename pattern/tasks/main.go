package main

import (
	"context"
	"log"
)

type Action interface {
	Do(ctx context.Context) error
}

type Tasks []Action

func (t Tasks) Do(ctx context.Context) error {
	for _, a := range t {
		if err := a.Do(ctx); err != nil {
			return err
		}
	}
	return nil
}

type Action1 struct {
}

func (a1 *Action1) Do(ctx context.Context) error {
	log.Print("this is action1")
	return nil
}

type ActionFunc func(ctx context.Context) error

// Do executes the func f using the provided context and frame handler.
func (f ActionFunc) Do(ctx context.Context) error {
	return f(ctx)
}

func Run(ctx context.Context, actions ...Action) error {
	//logic
	return Tasks(actions).Do(ctx)
}
func main() {
	ctx := context.Background()
	var actions []Action
	actions = append(actions, &Action1{})
	actions = append(actions, ActionFunc(func(ctx context.Context) error {
		log.Print("action func")
		return nil
	}))
	//Tasks(actions).Do(ctx)
	Run(ctx, Tasks{
		&Action1{},
		ActionFunc(func(ctx context.Context) error {
			log.Print("action fun")
			return nil
		}),
	})
}
