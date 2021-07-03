package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(linearSearch(arr, 7))
	fmt.Println(binarySearch(arr, 3))
}
