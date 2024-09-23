package main

import "fmt"

func main() {
	a := [5]int{}
	fmt.Print("Please enter the elements of the array:  ")

for i:=0;i<len(a);i++{
	fmt.Scanf("%d",&a[i])
}
fmt.Println(a)

}
