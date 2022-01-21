package main

import (
	"fmt"
	"log"
	"os"
)

func CreateFile() {
	file, err := os.Create("test.txt") // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close() // Make sure to close the file when you're done
	len, err := file.WriteString("Sudeb Dolui")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile Name: %s", file.Name())
}

func main() {
	fmt.Printf("########Create a file and Write the content #########\n")
	CreateFile()
}
