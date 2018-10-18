package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	n := 0
	for n >= 0 {
		n = <-c
		fmt.Print(n, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{2, 6, 7, 8, 3, 2, 6, 5, 3, 5, -1}

	go printCount(c)

	for _, v := range a {
		c <- v
	}

	time.Sleep(time.Millisecond * 1)
	fmt.Println("End")
}
