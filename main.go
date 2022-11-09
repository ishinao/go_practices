package main

func main() {
	linearTable([]int{90, 70, 50, 80, 60, 85})
	linearTableInsertOne([]int{90, 70, 50, 80, 60, 85}, 2, 75)
	linearTableAppendOne([]int{90, 70, 50, 80, 60, 85}, 75)
	linearTableRemoveOne([]int{90, 70, 50, 80, 60, 85}, 2)
	bubbleSort([]int{60, 50, 95, 80, 70})
}
