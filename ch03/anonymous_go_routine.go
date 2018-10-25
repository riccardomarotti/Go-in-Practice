package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("outside")

	go func() {
		fmt.Println("inside")
	}()

	fmt.Println("outside again")

	runtime.Gosched()
}
