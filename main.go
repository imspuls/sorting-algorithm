package main

import (
	"fmt"
	"sorting-algorithm/algorithm"
)

func main() {
	array := []int{6, 5, 3, 1, 8, 7, 2, 4}
	fmt.Println(algorithm.BubbleSort(array))
	fmt.Println(algorithm.InsertionSort(array))
	fmt.Println(algorithm.SelectionSort(array))
	fmt.Println(algorithm.ShellSort(array))
}
