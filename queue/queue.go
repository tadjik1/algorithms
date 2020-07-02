package queue

type Node struct {
	item interface{}
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	n    int
}

func New() *Queue {
	return &Queue{}
}

func (queue *Queue) IsEmpty() bool {
	return queue.n == 0
}

func (queue *Queue) Size() int {
	return queue.n
}

func (queue *Queue) Enqueue(item interface{}) {
	tail := queue.tail
	queue.tail = &Node{item, nil}
	if queue.IsEmpty() {
		queue.head = queue.tail
	} else {
		tail.next = queue.tail
	}
	queue.n = queue.n + 1
}

func (queue *Queue) Dequeue() interface{} {
	head := queue.head
	queue.head = head.next
	queue.n = queue.n - 1
	if queue.IsEmpty() {
		queue.tail = nil
	}
	return head.item
}

func (queue *Queue) Iterate() chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		node := queue.head
		for node != nil {
			ch <- node.item
			node = node.next
		}
	}()

	return ch
}
