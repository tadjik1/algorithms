package main

import "fmt"

/*
Deque, a double-ended queue is like a stack or a queue but
supports adding and removing items at both ends.
*/

type Node struct {
	item interface{}
	next *Node
	prev *Node
}

type Deque struct {
	left *Node
	right *Node
	n int
}

func New() *Deque {
	return &Deque{}
}

func (deque *Deque) IsEmpty() bool {
	return deque.n == 0
}

func (deque *Deque) Size() int {
	return deque.n
}

func (deque *Deque) PushLeft(item interface{}) {
	deque.left = &Node{item, deque.left, nil}
	if deque.IsEmpty() {
		deque.right = deque.left
	} else {
		deque.left.next.prev = deque.left
	}
	deque.n = deque.n + 1
}

func (deque *Deque) PushRight(item interface{}) {
	deque.right = &Node{item, nil, deque.right}
	if deque.IsEmpty() {
		deque.left = deque.right
	} else {
		deque.right.prev.next = deque.right
	}
	deque.n = deque.n + 1
}

func (deque *Deque) PopLeft() interface{} {
	left := deque.left
	deque.left = left.next
	deque.left.prev = nil
	deque.n = deque.n - 1
	if deque.IsEmpty() {
		deque.right = nil
	}
	return left.item
}

func (deque *Deque) PopRight() interface{} {
	right := deque.right
	right.next = nil
	deque.right = right.prev
	deque.n = deque.n - 1
	if deque.IsEmpty() {
		deque.left = nil
	}
	return right.item
}

func (deque *Deque) Iterate() chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		node := deque.left
		for node != nil {
			ch <- node.item
			node = node.next
		}
	}()

	return ch
}


func main() {
	d := New()
	values := []int{1, 2, 3, 4, 5}

	d.PushLeft(&values[0])
	d.PushLeft(&values[1])
	d.PushLeft(&values[2])
	d.PushLeft(&values[3])
	d.PushLeft(&values[4])

	for !d.IsEmpty() {
		fmt.Printf("value: %d, size: %d, isEmpty?: %t \n", *(d.PopRight()).(*int), d.Size(), d.IsEmpty())
	}

	d.PushRight(&values[0])
	d.PushRight(&values[1])
	d.PushRight(&values[2])
	d.PushRight(&values[3])
	d.PushRight(&values[4])

	for !d.IsEmpty() {
		fmt.Printf("value: %d, size: %d, isEmpty?: %t \n", *(d.PopLeft()).(*int), d.Size(), d.IsEmpty())
	}
}
