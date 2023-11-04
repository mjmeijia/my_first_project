package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		r float64 = 1.1
	)

	fmt.Println(area(r))

}

func area(r float64) float64 {
	Pi := math.Pi
	area := Pi * r * r
	return area
}
