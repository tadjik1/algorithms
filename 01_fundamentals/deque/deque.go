package deque

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
	left  *Node
	right *Node
	n     int
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
	node := &Node{item: item}
	if deque.IsEmpty() {
		deque.left = node
		deque.right = node
	} else {
		node.next = deque.left
		deque.left.prev = node
		deque.left = node
	}

	deque.n = deque.n + 1
}

func (deque *Deque) PushRight(item interface{}) {
	node := &Node{item: item}
	if deque.IsEmpty() {
		deque.right = node
		deque.left = node
	} else {
		node.prev = deque.right
		deque.right.next = node
		deque.right = node
	}

	deque.n = deque.n + 1
}

func (deque *Deque) PopLeft() interface{} {
	left := deque.left

	if left.next != nil {
		deque.left = left.next
		deque.left.prev = nil
	} else {
		deque.left = nil
		deque.right = nil
	}

	deque.n = deque.n - 1

	return left.item
}

func (deque *Deque) PopRight() interface{} {
	right := deque.right

	if right.prev != nil {
		deque.right = right.prev
		deque.right.next = nil
	} else {
		deque.right = nil
		deque.left = nil
	}

	deque.n = deque.n - 1

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
