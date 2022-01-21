package main

import "fmt"

func main() {

	// declaring an array of values
	array := []int{2, 4, 6, 8, 34, 5, 7, 3, 5, 67, 4}

	// size of the array
	n := 11

	// declaring a variable
	// to store the sum
	sum := 0

	// traversing through the
	// array using for loop
	for i := 0; i < n; i++ {

		// adding the values of
		// array to the variable sum
		sum += (array[i])
	}

	// declaring a variable
	// avg to find the average
	avg := (float64(sum)) / (float64(n))

	// typecast all values to float
	// to get the correct result
	fmt.Println("Sum = ", sum, "\nAverage = ", avg)
}
