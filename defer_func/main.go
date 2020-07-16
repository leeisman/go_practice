package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		func() {
			defer func(i int) {
				fmt.Println(i, "close")
			}(i)
			fmt.Println(i, "start")
		}()
	}

	for i := 4; i < 7; i++ {
		defer func(i int) {
			fmt.Println(i, "close")
		}(i)
		fmt.Println(i, "start")
	}
}
