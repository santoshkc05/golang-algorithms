package main

import "fmt"

func partition(slice []int) int {
	high := len(slice) - 1
	pivot := slice[high]
	//Temporary pivot index
	i := 0

	for j := 0; j < high; j++ {
		if slice[j] <= pivot {
			slice[i], slice[j] = slice[j], slice[i]
			i++
		}
	}

	// swap pivot value with temporary pivot index
	slice[i], slice[high] = slice[high], slice[i]
	return i
}

func quickSort(slice []int) {
	if len(slice) <= 1 {
		return
	}

	p := partition(slice)
	quickSort(slice[:p])
	quickSort(slice[p+1:])
}

func main() {
	slice := []int{50, 1, 2, 3, 3, 4, 10, 5, 7, 40, 29}
	quickSort(slice)
	fmt.Println(slice)
}
