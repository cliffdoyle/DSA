package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Top      int
	Capacity uint
	Array    []interface{}
}

// Return an initialized stack
func (stack *Stack) Init(Capacity uint) *Stack {
	stack.Top = -1
	stack.Capacity = Capacity
	stack.Array = make([]interface{}, Capacity)
	return stack
}

// Return a new stack
func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity)
}

// Stack is full when top is equal to the last index
func (stack *Stack) IsFull() bool {
	return stack.Top == int(stack.Capacity)-1
}

// Stack is empty when top is equal to -1
func (stack *Stack) IsEmpty() bool {
	return stack.Top == -1
}

// Returns the size of the size
func (stack *Stack) Size() uint {
	return uint(stack.Top + 1)
}

// Push adds data to stack
func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		return errors.New("stack is full")
	}
	stack.Top++
	stack.Array[stack.Top] = data
	return nil
}

// Pop removes data from the stack
func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	temp := stack.Array[stack.Top]
	stack.Top--
	return temp, nil
}

// Peek returns last inserted element without removing it
func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	temp := stack.Array[stack.Top]
	return temp, nil
}

// Drain removes all elements that are currently in the stack
func (stack *Stack) Drain() {
	stack.Array = nil
	stack.Top = -1
}

func main() {
	stack := NewStack(100)
	stack.Push(100000)
	fmt.Println("Pushed to stack : 10")
	stack.Push(20)
	fmt.Println("Pushed to stack : 20")
	stack.Push(30)
	fmt.Println("Pushed to stack : 30")
	data, _ := stack.Peek()
	fmt.Println("Top element : ", data)
	data, _ = stack.Pop()
	fmt.Println("Popped from stack : ", data)
	stack.Drain()
	fmt.Println("Stack size: ", stack.Size()) // prints 0 as stack size
}
