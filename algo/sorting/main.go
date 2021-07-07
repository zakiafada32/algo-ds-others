package main

import "fmt"

func main() {
	sort := []int{5, 4, 2, 7, 8, 9, 3, 6, 1}
	fmt.Println(sort)
	fmt.Println(merge(sort))
	fmt.Println(quick(sort))
}
