package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	//10 iterations
	// for i := 0; i < 10; i++ {
	// 	z -= (z*z - x) / (2 * z)
	// }
	// return z

	delta := 0.00000000001
	for {
		next := z - (z*z-x)/(2*z)
		if math.Abs(next-z) < delta {
			break
		}
		z = next
	}
	return z
}

func main() {
	for x := 1.0; x <= 10.0; x++ {
		fmt.Println("x:", x, "sqrt(x):", Sqrt(x))
	}
}
