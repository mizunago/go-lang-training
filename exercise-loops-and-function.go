package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z, old_z float64 = 1, 2
	for i := 0; i < 100000000; i++ {
		z = z - ((z*z - x) / float64(2*z))
		if old_z-z < 0.00000001 {
			fmt.Println(i)
			break
		} else {
			old_z = z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
