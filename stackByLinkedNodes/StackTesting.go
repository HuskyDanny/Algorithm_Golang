package main

import (
	stackByNodes "./Stack"
)

func main()  {
	stack := stackByNodes.StackInit()

	for i := 0; i <= 10; i++ {
		stack.Push(string(i+100))
	}
	stack.PrintAll()
	stack.Pop()
	stack.PrintAll()
	
}