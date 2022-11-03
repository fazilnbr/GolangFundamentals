package utility

import "fmt"

func getName() string {
	var name string
	fmt.Println("enter name : ")
	fmt.Scan(&name)
	return name
}
