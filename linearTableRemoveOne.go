package main

func linearTableRemoveOne(original []int, removePosition int) []int {
	result := append(original[:removePosition], original[removePosition+1:]...)
	return result
}
