package main

import (
	"fmt"

	ds "github.com/vagnerclementino/algoritmhs/go/ds/linkedlist"
)

func main() {
	linkedList := ds.NewLinkedList()
	fmt.Println(linkedList)
	for i := range 100 {
		linkedList.Add(i)
	}
	fmt.Println(linkedList)
}
