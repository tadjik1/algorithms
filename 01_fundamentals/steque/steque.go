package steque

import "fmt"

/*
A stack-ended queue (steque) is a data type that supports push, pop,
and enqueue.
*/

type Node struct {
	item interface{}
	next *Node
}

type Steque struct {
	head *Node
	tail *Node
	n    int
}

func New() *Steque {
	return &Steque{}
}

func (steque *Steque) IsEmpty() bool {
	return steque.n == 0
}

func (steque *Steque) Size() int {
	return steque.n
}

func (steque *Steque) Push(item interface{}) {
	steque.head = &Node{item, steque.head}
	if steque.n == 0 {
		steque.tail = steque.head
	}
	steque.n = steque.n + 1
}

func (steque *Steque) Pop() interface{} {
	head := steque.head
	steque.head = head.next
	steque.n = steque.n - 1
	if steque.IsEmpty() {
		steque.tail = nil
	}
	return head.item
}

func (steque *Steque) Enqueue(item interface{}) {
	tail := steque.tail
	steque.tail = &Node{item, nil}
	if steque.IsEmpty() {
		steque.head = steque.tail
	} else {
		tail.next = steque.tail
	}
	steque.n = steque.n + 1
}

func (steque *Steque) Iterate() chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		node := steque.head
		for node != nil {
			ch <- node.item
			node = node.next
		}
	}()

	return ch
}

func main() {
	s := New()
	values := []int{1, 2, 3, 4, 5}

	s.Push(&values[0])
	s.Push(&values[1])
	s.Push(&values[2])
	s.Push(&values[3])
	s.Push(&values[4])

	for !s.IsEmpty() {
		fmt.Printf("value: %d, size: %d, isEmpty?: %t \n", *(s.Pop()).(*int), s.Size(), s.IsEmpty())
	}

	s.Enqueue(&values[0])
	s.Enqueue(&values[1])
	s.Enqueue(&values[2])
	s.Enqueue(&values[3])
	s.Enqueue(&values[4])

	for !s.IsEmpty() {
		fmt.Printf("value: %d, size: %d, isEmpty?: %t \n", *(s.Pop()).(*int), s.Size(), s.IsEmpty())
	}
}
