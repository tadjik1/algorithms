package main

import (
	"algorithms/bag"
	"algorithms/queue"
	"algorithms/stack"
	"fmt"
)

func main() {
	fmt.Println("stack")
	s := stack.New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	for !s.IsEmpty() {
		fmt.Printf("value: %d, size: %d, isEmpty?: %t \n", *s.Pop(), s.Size(), s.IsEmpty())
	}

	fmt.Println("queue")
	q := queue.New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	for !q.IsEmpty() {
		fmt.Printf("value: %d, size: %d, isEmpty?: %t \n", *q.Dequeue(), q.Size(), q.IsEmpty())
	}

	fmt.Println("bag")
	b := bag.New()
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	b.Add(5)

	for value := range b.Iterate() {
		fmt.Printf("value: %d, size: %d, isEmpty?: %t \n", *value, b.Size(), b.IsEmpty())
	}
}
