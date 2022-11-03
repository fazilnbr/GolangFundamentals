package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println(time.Now(), "taking a nap")
		fmt.Println("enter")
		var a int
		fmt.Scan(&a)
		ch <- a

		time.Sleep(2 * time.Second)

	}()

	fmt.Println(time.Now(), "waiting for message")

	v := <-ch

	fmt.Println(time.Now(), "received", v)
}
