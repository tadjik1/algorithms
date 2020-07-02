package bag

type Node struct {
	item interface{}
	next *Node
}

type Bag struct {
	head *Node
	n    int
}

func New() *Bag {
	return &Bag{}
}

func (stack *Bag) IsEmpty() bool {
	return stack.n == 0
}

func (stack *Bag) Size() int {
	return stack.n
}

func (stack *Bag) Add(item interface{}) {
	stack.head = &Node{item, stack.head}
	stack.n = stack.n + 1
}

func (stack *Bag) Iterate() chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		node := stack.head
		for node != nil {
			ch <- node.item
			node = node.next
		}
	}()

	return ch
}
