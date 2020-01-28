package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next *Node
}

var size = 0
var queue = new(Node)

func push(t *Node, v int) bool {
	if queue == nil {
		queue = &Node{v,nil}
		size ++
		return true
	}
	t = &Node{v,nil}
	t.Next = queue
	queue = t
	size ++
	return true
}

func pop(t *Node) (int, bool) {
	if size == 0 {
		return 0,false
	}

	if size == 1 {
		queue = nil
		size --
		return t.Value, true
	}

	temp := t

	for (t.Next) != nil {
		temp = t
		t = t.Next
	}

	v := (temp.Next).Value
	temp.Next = nil

	size -- 
	return v,true
}

func traverse(t *Node) {
	if size == 0 {
		fmt.Println("Empty Queue!")
		return
	}

	for t != nil {
		fmt.Printf("%d ->", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	queue = nil
	push(queue,10)
	fmt.Println("Size : ", size)
	traverse(queue)

	v, b := pop(queue)
	if b {
		fmt.Println("pop : ", v)
	}
	fmt.Println("Size : ", size)

	for i := 0; i < 5; i ++ {
		push(queue, i)
	}

	traverse(queue)
	fmt.Println("Size : ", size)

	v, b = pop(queue)

	if b {
		fmt.Println("pop : ", v)
	}
	fmt.Println("Size : ", size)

	v, b = pop(queue)
	if b {
		fmt.Println("pop : ", v)
	}
	fmt.Println("Size : ", size)
	traverse(queue)
}