package main

import "fmt"

type node struct {
	data int
	next *node
}
type Queue struct {
	Frond *node
	Rear  *node
}

func (list *Queue) EnQueue() {
	var num int
	fmt.Println("Enter the number :")
	fmt.Scan(&num)
	NewQueue := new(node)
	NewQueue.data = num
	NewQueue.next = nil

	if list.Frond == nil {
		list.Frond = NewQueue
	} else {
		list.Rear.next = NewQueue
	}
	list.Rear = NewQueue
}

func (list *Queue) DeQueue() {
	if list.Frond == nil {
		fmt.Println("Queue is NULL!")
	} else {
		list.Frond = list.Frond.next
	}

}

func (list *Queue) Display() {
	fmt.Print("\nThe Queue is : ")
	temp := list.Frond
	for temp != nil {
		fmt.Print(" ", temp.data)
		temp = temp.next
	}

}

func (list *Queue) get() {
loop:
	for {

		fmt.Println("\n\nThis program for singlelinkedlist\n1-Insert\n2-Display\n3-Delete\n4-Exit\n\nChoose Options  :")
		var op int
		fmt.Scan(&op)
		fmt.Println()
		switch op {
		case 1:
			list.EnQueue()
		case 2:
			list.Display()
		case 3:
			list.DeQueue()
		case 4:
			break loop
		default:
			fmt.Println("Choose correct input")
		}

	}
}

func main() {

	list := Queue{}

	list.get()

}
