package main

import (
	"fmt"
	"math/rand"
	"time"

	ds "github.com/vagnerclementino/algoritmhs/go/ds/linkedlist"

	graph "github.com/vagnerclementino/algoritmhs/go/ds/graph"
	heap "github.com/vagnerclementino/algoritmhs/go/ds/priorityqueue"
	queue "github.com/vagnerclementino/algoritmhs/go/ds/queue"
	set "github.com/vagnerclementino/algoritmhs/go/ds/set"
	stack "github.com/vagnerclementino/algoritmhs/go/ds/stack"
	tree "github.com/vagnerclementino/algoritmhs/go/ds/tree"
)

func main() {

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

	newSet.Add(0)
	fmt.Println(newSet)

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

	newStack := stack.NewStack()
	for i := range 5 {
		newStack.Push(i)
	}
	fmt.Println(newStack)

	for !newStack.IsEmpty() {
		fmt.Println(newStack.Pop())
	}

	rand.Seed(time.Now().UnixNano())
	bst := tree.NewBinarySearchTree()

	for i := 0; i < 10; i++ {
		bst.Add(rand.Intn(100))
	}
	fmt.Println(bst)

	if min, err := bst.MinValue(); err == nil {
		fmt.Printf("the min value for current tree is %d\n", min)
	}

	if max, err := bst.MaxValue(); err == nil {
		fmt.Printf("the max value for current tree is %d\n", max)
	}
	maxHeap := heap.NewMaxHeap()
	for i := 0; i < 10; i++ {
		val := randomRange(-100, 100)
		maxHeap.Add(val)
		fmt.Printf("adding the value %d at heap\n", val)
	}

	if max, ok := maxHeap.Peak(); ok {
		fmt.Printf("The Max value in the heap is %d\n", max)
	}

	minHeap := heap.NewMinHeap()
	for i := 0; i < 10; i++ {
		val := randomRange(-100, 100)
		minHeap.Add(val)
		fmt.Printf("adding the value %d at heap\n", val)
	}

	if min, ok := minHeap.Peak(); ok {
		fmt.Printf("The min value in the heap is %d\n", min)
	}

	if max, ok := maxHeap.Peak(); ok {
		fmt.Printf("The Max value in the heap is %d\n", max)
	}

	for i := 0; i < 10; i++ {
		val := randomRange(-100, 100)
		minHeap.Add(val)
		fmt.Printf("adding the value %d at heap\n", val)
	}

	if min, ok := minHeap.Peak(); ok {
		fmt.Printf("The min value in the heap is %d\n", min)
	}

	for !maxHeap.IsEmpty() {
		if max, ok := maxHeap.Pop(); ok {
			fmt.Printf("Removed the max value %d from the max heap\n", max)
		}
	}

	// for !minHeap.IsEmpty() {
	// 	if min, ok := minHeap.Pop(); ok {
	// 		fmt.Printf("Removed the min value %d from the min heap", min)
	// 	}
	// }
	g := graph.NewGraph(true)
	g.AddEdge("A", "B", 1.0)
	g.AddEdge("A", "C", 1.0)
	g.AddEdge("B", "D", 1.0)
	g.AddEdge("C", "D", 1.0)
	g.AddEdge("D", "E", 1.0)

	fmt.Println("BFS a partir de A:")
	graph.BFS(g, "A", func(n graph.Node) {
		fmt.Println(n)
	})
}

func randomRange(min, max int) int {
	return min + rand.Intn(max-min+1)
}
