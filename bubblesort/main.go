package main

import (
	"fmt"
	"math/rand"
)

func makeRandomSlice(numItems, max int) []int {
	slices := make([]int, numItems)
	for i := range slices {
		slices[i] = rand.Intn(max) + 1
	}
	return slices
}

func printSlice(slice []int, numItems int) {
	itemsToPrint := min(len(slice), numItems)
	fmt.Println(slice[:itemsToPrint])
}

func checkSorted(slice []int) {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			fmt.Println("Not sorted")
			return
		}
	}
	fmt.Println("sorted")
}

func bubbleSort(slice []int) {
	for {
		swapped := false
		n := len(slice)

		for i := 1; i < n; i++ {
			if slice[i-1] > slice[i] {
				slice[i], slice[i-1] = slice[i-1], slice[i]
				swapped = true
			}
		}
		//n-- reduces the range of elements to compare after each pass, as the largest element of the unsorted part is sorted.
		n--
		if !swapped {
			break
		}
	}

	//fmt.Printf("Sorted: %v", slice)
}

func main() {
	// Get the number of items and maximum item value.
	var numItems = 10
	var max = 20

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}
