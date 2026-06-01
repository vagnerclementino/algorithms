package ds

import "fmt"

type DoublyNode struct {
	Value int
	Next  *DoublyNode
	Prev  *DoublyNode
}

type DoublyLinkedList struct {
	Head *DoublyNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) IsEmpty() bool {
	return l != nil && l.Head == nil
}

func (l *DoublyLinkedList) Add(item int) {
	newNode := &DoublyNode{Value: item}
	if l.IsEmpty() {
		l.Head = newNode
		return
	}

	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode
	newNode.Prev = currentNode
}

func (ll *DoublyLinkedList) String() string {
	if ll.IsEmpty() {
		return "(nil)"
	}
	var output string
	for node := ll.Head; node != nil; node = node.Next {
		output += fmt.Sprintf(" %d ->", node.Value)
	}
	output += " (nil)"
	return output
}
