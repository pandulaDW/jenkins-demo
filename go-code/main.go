package main

import "fmt"

func add(a, b int) int {
	return a+b
}

func main() {
	fmt.Println("Running go code..")
	result := add(4, 10)
	fmt.Println("result is:", result)
	fmt.Println("Ending go code..")
}