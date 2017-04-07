package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var z, old_z float64 = 1, 2
	for i := 0; i < 100000000; i++ {
		z = z - ((z*z - x) / float64(2*z))
		if old_z-z < 0.00000001 {
			break
		} else {
			old_z = z
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
