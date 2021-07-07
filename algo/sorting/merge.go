package main

func merge(array []int) []int {
	if len(array) == 1 {
		return array
	}

	middle := len(array) / 2
	left := array[:middle]
	right := array[middle:]

	array = mergeSort(merge(left), merge(right))

	return array
}

func mergeSort(left, right []int) []int {
	totalLenght := len(left) + len(right)
	response := make([]int, 0, totalLenght)
	lIndex, rIndex := 0, 0

	for lIndex < len(left) || rIndex < len(right) {
		if lIndex < len(left) && (rIndex == len(right) || left[lIndex] < right[rIndex]) {
			response = append(response, left[lIndex])
			lIndex++
		} else {
			response = append(response, right[rIndex])
			rIndex++
		}
	}

	return response
}
