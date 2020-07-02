package main

import (
	"algorithms/bag"
	"algorithms/queue"
	"algorithms/stack"
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4, 5}

	fmt.Println("stack")

	s := stack.New()
	s.Push(&values[0])
	s.Push(&values[1])
	s.Push(&values[2])
	s.Push(&values[3])
	s.Push(&values[4])

	for !s.IsEmpty() {
		fmt.Printf("value: %d, size: %d, isEmpty?: %t \n", s.Pop(), s.Size(), s.IsEmpty())
	}

	fmt.Println("queue")
	q := queue.New()
	q.Enqueue(&values[0])
	q.Enqueue(&values[1])
	q.Enqueue(&values[2])
	q.Enqueue(&values[3])
	q.Enqueue(&values[4])

	for !q.IsEmpty() {
		fmt.Printf("value: %d, size: %d, isEmpty?: %t \n", q.Dequeue(), q.Size(), q.IsEmpty())
	}

	fmt.Println("bag")
	b := bag.New()
	b.Add(&values[0])
	b.Add(&values[1])
	b.Add(&values[2])
	b.Add(&values[3])
	b.Add(&values[4])

	for value := range b.Iterate() {
		fmt.Printf("value: %d, size: %d, isEmpty?: %t \n", value, b.Size(), b.IsEmpty())
	}
}
