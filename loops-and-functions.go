package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, y := x, 0.0
	for math.Abs(y-z) > 1e-6 {
		y, z = z, z-(z*z-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println("Sqrt()'s result: ", Sqrt(2))
	fmt.Println("math.Sqrt()'s result: ", math.Sqrt(2))
}
