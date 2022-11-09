package main

func linearTable(original []int) []int {
	result := make([]int, len(original))
	for i := 0; i < len(original); i++ {
		result[i] = original[i]
	}
	return result
}
