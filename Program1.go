package main

import "fmt"

func main() {
	name := "venu" // dynamic type of variable according to the value stored
	fmt.Println(name)
	// define i
	var i int // statically mentioning data type
	// set value for i
	i = 10
	// we can also set value as follows
	var y = 5 // style coming from JavaScript
	var msg = "Remote host found."
	var foo, bar int = 100, 200
	fmt.Println(bar)
	vehicle := "creta"
	age := 62
	var is_job_failed = false
	fmt.Printf("%d %d %s\n", i, y, msg)
	fmt.Println(foo)
	fmt.Println(age)
	fmt.Println(vehicle)
	fmt.Println(is_job_failed)
}
