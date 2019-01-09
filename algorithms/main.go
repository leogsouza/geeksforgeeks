package main

import (
	"fmt"

	js "github.com/leogsouza/geeksforgeeks/algorithms/jumpsearch"
)

func main() {
	/*slc1 := []int{1, 3, 5, 4, 7, 9}
	n1 := 7
	fmt.Println("Find number at index ", ls.linearSearch(slc1, n1))

	slc2 := []int{10, 23, 15, 44, 37, 97, 48}
	n2 := 98
	fmt.Println("Find number at index ", ls.linearSearch(slc2, n2))*/

	// arr1 := []int{1, 5, 7, 9, 12}
	// totalArr1 := len(arr1)
	// n1 := 9
	// fmt.Println("Number found at index ", bs.BinarySearch(arr1, 0, totalArr1-1, n1))

	// arr2 := []int{2, 4, 5, 7, 12, 16, 18, 36}
	// length := len(arr2)
	// n2 := 18
	// fmt.Println("Number found at index ", bs.BinarySearchIterative(arr2, 0, length-1, n2))

	// arr := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	arr := []int{0, 1, 1, 2, 3, 5}

	n1 := 4
	index := js.JumpSearch(arr, n1)
	fmt.Printf("Number %d found at index %d", n1, index)

}
