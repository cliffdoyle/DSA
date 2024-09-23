package main

import "fmt"

type Edges struct {
	A int
	B int
}

func (e *Edges) Add() {
	e.A++
	e.B--
	fmt.Println(e.A,e.B)
}

func main() {
	// var m = map[string]string{
	// 	"kely": "right",
	// }

	// v, ok := m["kely"]
	// fmt.Println(v,ok)

	// j:=Edges{
	// 	A: 23,
	// 	B: 65,
	// }
	// j.Add()
	// j.Add()
	// j.Add()



	ch:=make(chan int,56)

	for i:=range ch{
		fmt.Println(i)
	}
}
