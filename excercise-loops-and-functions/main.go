package main

import (
	"fmt"
	"math"
)

const e = 1e-8

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		prev := z
		z -= (z*z - x) / (2*z)

		if math.Abs(prev - z) < e {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

