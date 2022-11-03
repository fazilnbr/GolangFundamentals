package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	s := time.Now()
	w := new(sync.WaitGroup)
	w.Add(4)
	go timeset(20000, w)
	go timeset(40000, w)
	go timeset(60000, w)
	go timeset(80000, w)
	w.Wait()
	a := time.Since(s)
	fmt.Println(a)
}

func timeset(t int, w *sync.WaitGroup) {
	time.Sleep(time.Duration(t))
	fmt.Println("hello", t)
	defer w.Done()
}
