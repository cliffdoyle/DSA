package main

import "fmt"

func main() {
	var a [10]int
	var size int
	var pos int

	fmt.Print("Enter size of the array: "+"\n")
	fmt.Scanf("%d",&size)
	fmt.Print("Which position do you want to delete: \n")
	fmt.Scan(&pos)

	fmt.Print("Enter data here: " + "\n")

	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &a[i])
	}

	fmt.Println(a)
	
	// pos := 3
	// num := 80
	// k:=34
	// z:=0

	//Inserting data in a sorted array by shifting all the elements to the right
	// for i := size; i >= pos-1; i-- {
	// 	a[i+1] = a[i]
	// }
	// a[pos-1]=num
	// fmt.Println("sorted array:",a)

	//Inserting data in an unsorted array
	// a[size]=a[pos-1]
	// a[pos-1]=k
	// fmt.Println("unsorted array:",a)

	//Deletion of a value from an array
	for i:=pos-1;i<size;i++{
		a[i]=a[i+1]
		// size--
	}
	fmt.Println("Deletion of an array: ",a)

	// fmt.Println(len(a))
}
