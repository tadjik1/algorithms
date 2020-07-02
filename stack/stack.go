package stack

type Node struct {
	item interface{}
	next *Node
}

type Stack struct {
	head *Node
	n    int
}

func New() *Stack {
	return &Stack{}
}

func (stack *Stack) IsEmpty() bool {
	return stack.n == 0
}

func (stack *Stack) Size() int {
	return stack.n
}

func (stack *Stack) Push(item interface{}) {
	stack.head = &Node{item, stack.head}
	stack.n = stack.n + 1
}

func (stack *Stack) Pop() interface{} {
	head := stack.head
	stack.head = head.next
	stack.n = stack.n - 1
	return head.item
}

func (stack *Stack) Peek() interface{} {
	return stack.head.item
}

func (stack *Stack) Iterate() chan interface{} {
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
