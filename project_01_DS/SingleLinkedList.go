package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}
type SingleLinkedList struct {
	Head *node
	Tail *node
}

func (list *SingleLinkedList) AddNode(data int) {
	newNode := new(node)
	newNode.data = data
	newNode.next = nil

	if list.Head == nil {
		list.Head = newNode
	} else {
		list.Tail.next = newNode
	}
	list.Tail = newNode
	// fmt.Println("node added")
}

func (list *SingleLinkedList) display() {
	fmt.Print("The Linked List is : ")
	if list.Head == nil {
		fmt.Println("list is empty")
	}
	temp := list.Head

	for temp != nil {
		fmt.Print(" ", temp.data)
		temp = temp.next
	}
}

func (list *SingleLinkedList) Delete(data int) {
	temp := list.Head

	for temp != nil {
		if temp.next.data == data {
			temp.next = temp.next.next
			break
		}
		temp = temp.next
	}
}

func (list *SingleLinkedList) get() {
loop:
	for {

		fmt.Println("\n\nThis program for singlelinkedlist\n1-Insert\n2-Display\n3-Delete\n4-Exit\n\nChoose Options  :")
		var op, num int
		fmt.Scan(&op)
		fmt.Println()
		switch op {
		case 1:
			fmt.Println("Enter the number :")
			fmt.Scan(&num)
			list.AddNode(num)
		case 2:
			list.display()
		case 3:
			fmt.Println("Enter the number :")
			fmt.Scan(&num)
			list.Delete(num)
		case 4:
			break loop
		default:
			fmt.Println("Choose correct input")
		}

	}
}

func main() {
	list := SingleLinkedList{}

	list.get()

}
