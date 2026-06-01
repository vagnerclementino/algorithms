package main

import (
	"fmt"

	ds "github.com/vagnerclementino/algoritmhs/go/ds/linkedlist"
	queue "github.com/vagnerclementino/algoritmhs/go/ds/queue"
	set "github.com/vagnerclementino/algoritmhs/go/ds/set"
)

func main() {

	fmt.Println("---LINKED LIST---")
	linkedList := ds.NewLinkedList()
	for i := range 10 {
		linkedList.Add(i)
		fmt.Println(linkedList)
	}

	fmt.Println("--------")

	fmt.Println("---DOUBLYLINKED LIST---")
	doublyLinkedList := ds.NewDoublyLinkedList()
	for i := range 10 {
		doublyLinkedList.Add(i)
		fmt.Println(doublyLinkedList)
	}

	fmt.Println("---SET---")

	newSet := set.NewSet()
	for i := range 10 {
		newSet.Add(i)
	}

	setEven := set.NewSet()
	for i := range 10 {

		if i%2 == 0 {
			setEven.Add(i)
		}
	}
	fmt.Println(newSet.Intersect(setEven))

	fmt.Println("---SET---")
	defaultQueue := queue.NewQueue()

	for i := range 10 {
		defaultQueue.Enqueue(i)
	}

	fmt.Println(defaultQueue)
	if err, _ := defaultQueue.Dequeue(); err == nil {
		fmt.Println("dequeue with success")
	}
	fmt.Println(defaultQueue)

	sortedQueue := queue.NewSortedQueue()

	for i := range 10 {
		sortedQueue.Add(i)
	}

	fmt.Println(sortedQueue)
}
