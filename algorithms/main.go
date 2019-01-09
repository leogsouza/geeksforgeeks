package main

import (
	"fmt"

	ls "github.com/leogsouza/geeksforgeeks/algorithms/linearsearch"
)

func main() {
	slc1 := []int{1, 3, 5, 4, 7, 9}
	n1 := 7
	fmt.Println("Find number at index ", ls.linearSearch(slc1, n1))

	slc2 := []int{10, 23, 15, 44, 37, 97, 48}
	n2 := 98
	fmt.Println("Find number at index ", ls.linearSearch(slc2, n2))
}
