package main

import (
	"algorithms/stack"
	"fmt"
	"strconv"
)

/*
Dijkstra's Two-Stack Algorithm for Expression Evaluation
*/

func main() {
	ops := stack.New()
	vals := stack.New()

	for _, c := range "(1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )" {
		s := string(c)
		switch s {
		case "(":
		case " ":
			continue
		case "+":
			ops.Push(s)
		case "*":
			ops.Push(s)
		case ")":
			val1 := (vals.Pop()).(int)
			val2 := (vals.Pop()).(int)
			op := ops.Pop()

			switch op {
			case "+":
				vals.Push(val1 + val2)
			case "*":
				vals.Push(val1 * val2)
			}

		default:
			val, _ := strconv.Atoi(s)
			vals.Push(val)
		}
	}

	fmt.Println(vals.Pop())
}
