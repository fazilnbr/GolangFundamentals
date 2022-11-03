package main

import "fmt"

type node struct {
	data int
	next *node
}
type stack struct {
	Top *node
}

func (list *stack) push() {
	var num int
	fmt.Println("Enter the number :")
	fmt.Scan(&num)
	NewStack := new(node)
	NewStack.data = num
	NewStack.next = nil

	if list.Top == nil {
		list.Top = NewStack
	} else {
		list.Top.next = NewStack
	}

	list.Top = NewStack
}

func (list *stack) pop() {
	list.Top = list.Top.next
}

func (list *stack) Display() {
	temp := list.Top
	fmt.Println("the stack : ")
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
}

func (list *stack) get() {
loop:
	for {

		fmt.Println("\n\nThis program for singlelinkedlist\n1-Insert\n2-Display\n3-Delete\n4-Exit\n\nChoose Options  :")
		var op int
		fmt.Scan(&op)
		fmt.Println()
		switch op {
		case 1:
			list.push()
		case 2:
			list.Display()
		case 3:
			list.pop()
		case 4:
			break loop
		default:
			fmt.Println("Choose correct input")
		}

	}
}

func main() {

	list := stack{}

	list.get()

}
