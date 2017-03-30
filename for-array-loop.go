package main

import (
	"fmt"
	"math"
)

var pow []int

func main() {
	for i := 0; i < 20; i++ {
		// tmp := math.Pow(2, float64(i)) // 冪乗計算は float64 への変換が必要
		tmp := math.Exp2(float64(i)) // 2 の冪乗は Exp2, 10 は Exp が使える
		pow = append(pow, int(tmp))
	}
	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value)
	}
}
