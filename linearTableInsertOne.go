package main

func linearTableInsertOne(original []int, insertValue int, insertPosition int) []int {
	result := append(original[:insertPosition+1], original[insertPosition:]...)
	result[insertPosition] = insertValue
	return result
}
