package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

//Notice that we changed the fmt.Println to be a return instead. The return statement causes the function to immediately stop and return the value after it to the function that called this one. Modify main to look like this:
func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
}
