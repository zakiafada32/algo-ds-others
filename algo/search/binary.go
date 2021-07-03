package main

func binarySearch(arr []int, key int) bool {
	if len(arr) == 1 {
		return arr[0] == key
	}

	middle := len(arr) / 2
	if arr[middle] == key {
		return true
	} else if arr[middle] > key {
		return binarySearch(arr[:middle], key)
	} else {
		return binarySearch(arr[middle+1:], key)
	}
}
