package main

import "fmt"

type Node struct {
	Prev *Node
	Data int
	Next *Node
}

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
}

func (list *DoubleLinkedList) AddNode(data int) {
	newNode := new(Node)
	newNode.Data = data
	newNode.Prev = nil
	newNode.Next = nil

	if list.Head == nil {
		list.Head = newNode

	} else {
		newNode.Prev = list.Tail
		list.Tail.Next = newNode
	}
	list.Tail = newNode
}
func (list *DoubleLinkedList) display() {
	fmt.Print("\nThe list are :")
	if list.Head == nil {
		fmt.Println("List is empty ")
	}
	temp := list.Head
	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}

}

func (list *DoubleLinkedList) displayreverce() {
	fmt.Print("\nThe list are :")
	if list.Head == nil {
		fmt.Println("List is empty ")
	}
	temp := list.Tail
	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Prev
	}

}

func (list *DoubleLinkedList) delete() {
	fmt.Println("enter the number :")
	var num int
	fmt.Scan(&num)
	temp := list.Head
	for temp != nil {
		if temp.Data == num {
			if temp == list.Head {
				list.Head = list.Head.Next
				list.Head.Prev = nil
			} else if list.Tail == temp {
				list.Tail = temp.Prev
				list.Tail.Next = nil
			} else {
				temp.Next.Prev = temp.Prev
				temp.Prev.Next = temp.Next
				break
			}
		}
		temp = temp.Next
	}

}

func (list *DoubleLinkedList) get() {
loop:
	for {

		fmt.Println("\n\nThis program for Doublelinkedlist\n1-Insert\n2-Display\n3-Display Reverce\n4-Delete\n5-Exit\n\nChoose Options  :")
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
			list.displayreverce()
		case 4:
			list.delete()
		case 5:
			break loop
		default:
			fmt.Println("Choose correct input")
		}

	}
}

func main() {
	list := DoubleLinkedList{}
	list.get()

}
