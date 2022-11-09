package main

func bubbleSort(original []int) []int {
	result := bubbleSortFunc(original)
	return result
}

func bubbleSortFunc(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		isSwapped := false
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
				isSwapped = true
			}
		}
		if !isSwapped {
			break
		}
	}
	return array
}
