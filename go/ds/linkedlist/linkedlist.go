package ds

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) IsEmpty() bool {
	return l != nil && l.Head == nil
}

func (l *LinkedList) Add(item int) {
	newNode := &Node{Value: item}
	if l.IsEmpty() {
		l.Head = newNode
		return
	}
	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode
}

func (ll *LinkedList) String() string {
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
