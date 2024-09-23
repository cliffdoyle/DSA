package main

import (
	"fmt"
	"strconv"
)

func main() {
	// arr := []int{3, 9, 5, 9, 7, 1}

	// for i := len(arr)-1; i >=0; i-- {
	// 	for j := len(arr)-1; j>=i; j-- {
	// 		if arr[i] > arr[j] {
	// 			arr[i], arr[j] = arr[j], arr[i]
	// 		}
	// 	}
	// }

	// fmt.Println(arr)
	arr := []int{5, 7, 3, 9, 1, 9}
	greedyAlgo(arr)
}

func greedyAlgo(arr []int) {
	// arr=[]int{5,7,3,9,1,9}
	num := []int{}
	for len(arr)>0{
	m := 0
	for i:=1;i<len(arr);i++{
		if arr[i]>arr[m]{
			m=i
		}
	}
	num = append(num, arr[m])
	arr = append(arr[:m],arr[m+1:]... )
}
	res:=""
	for _,val:=range num{
		res+=strconv.Itoa(val)
	}
	// m largest
	fmt.Println(res)
	fmt.Println(num)
}
