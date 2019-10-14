package main

import (
	"fmt"
)

type Node struct {
	content int
	next    *Node
}

type ItemLinkedList struct {
	head *Node
	size int
	//lock sync.RWMutex
}

func (ll *ItemLinkedList) Append(t int) {
	node := &Node{t, nil}
	if ll.head == nil {
		ll.head = node
	} else {
		currentPosition := ll.head
		for currentPosition.next != nil {
			currentPosition = currentPosition.next
		}
		currentPosition.next = node
	}
	ll.size++
}

func (ll *ItemLinkedList) Insert(i, t int) error {
	if i < 0 || i > ll.size {
		return fmt.Errorf("Index Out of Bound")
	}
	newNode := &Node{t, nil}

	if i == 0 {
		newNode.next = ll.head
		ll.head = newNode
		return nil
	}

	currentPosition := ll.head
	j := 0
	for j < i-1 && currentPosition.next != nil {
		j++
		currentPosition = currentPosition.next

	}
	prev := currentPosition
	nextNode := currentPosition.next

	prev.next = newNode
	newNode.next = nextNode

	return nil
}

func (ll *ItemLinkedList) String() {
	currentPosition := ll.head
	for currentPosition != nil {
		fmt.Print(currentPosition.content)
		fmt.Print(" ")
		currentPosition = currentPosition.next
	}
	fmt.Println()
}

func (ll *ItemLinkedList) NthLastElement(n int) int {
	p, f := ll.head, ll.head

	if n < 0 || n > ll.size {
		return -1
	}

	i := 0
	for i <= n && f != nil {
		i++
		f = f.next
	}

	for f != nil && p.next != nil {
		p = p.next
		f = f.next
	}

	return p.content
}

func (ll *ItemLinkedList) PrintReverseList() {
	printReversedList(ll.head)
}

func printReversedList(node *Node) {
	if node == nil {
		return
	}
	printReversedList(node.next)
	fmt.Printf("%d ", node.content)
}

func main() {
	ll := &ItemLinkedList{}

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)
	ll.String()

	ll.Insert(0, 90)
	ll.String()
	//fmt.Println(ll.NthLastElement(9))
	//ll.PrintReverseList()
}
