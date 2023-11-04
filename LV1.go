package main

import "fmt"

func main() {
	var (
		a, b int = 1, 2
	)

	c := add(a, b)
	fmt.Println(c)
}

func add(a, b int) int {
	return a + b
}
