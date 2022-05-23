package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func add(list *Node, data int) *Node {
	if list == nil {
		list := new(Node)
		list.value = data
		list.next = nil
		return list
	} else {
		temp := list

		for temp.next != nil {
			temp = temp.next
		}

		nn := new(Node)
		nn.value = data
		nn.next = nil

		temp.next = nn
		return list
	}
}

func display(list *Node) {
	var temp *Node
	temp = list
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}
func main() {
	var head *Node
	head = nil
	// add 5 data
	fmt.Println("=== insert 5 data=== ")
	n := 0
	for n < 5 {
		fmt.Printf("data %d\n", n)
		head = add(head, n)
		n++
	}
	fmt.Println("=== display ===")
	display(head)
}
