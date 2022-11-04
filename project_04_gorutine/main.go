package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	greetings("hello world")

	greetings("hi fasil")

	fmt.Println(time.Since(t))
}

func greetings(s string) {
	defer fmt.Println("end")
	for i := 0; i < 6; i++ {
		fmt.Println(s)

		time.Sleep(time.Second * 1)
	}
}
