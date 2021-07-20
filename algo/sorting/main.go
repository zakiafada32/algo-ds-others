package main

import "fmt"

func main() {
	sort := []int{9, 4, 2, 7, 8, 5, 3, 6, 1}
	fmt.Println(sort)
	fmt.Println(merge(sort))
	fmt.Println(quick(sort))
}
