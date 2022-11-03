package main

import (
	"fmt"
	"time"
)

func variadic(num ...int) {
	f := 1
	if num != nil {
		for _, n := range num {
			f *= n
		}
		fmt.Println(f)
	}

}
func main() {
	variadic()
	t := time.Now()

	// print location and local time
	fmt.Println(t.Location())
}
