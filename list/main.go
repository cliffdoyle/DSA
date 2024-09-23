package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	Head *Node
	Size int
}

func main() {
	// firs := &Node{Data: 67}
	// second := &Node{Data: 89}

	// firs.Next = second
	var l List
	l.AddData(67)
	l.AddData(10)
	l.AddData(124)
	// if l.Head == nil {
	// 	l.Head = firs
	// }
	l.Print()
	fmt.Println(l.length())
}

func (l *List) AddData(data int) {
	newNode := &Node{Data: data}
	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *List) Print() {
	current := l.Head

	for current != nil {
		fmt.Printf("%d->", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func (l *List) length() int {
	l.Size = 0

	current := l.Head
	for current != nil {
		l.Size++
		current = current.Next
	}


	return l.Size
}
