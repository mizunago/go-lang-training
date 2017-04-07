package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (a [][]uint8) {
	// return で返す変数を決めちゃう形式に変更
	a = make([][]uint8, dy)
	for i := range a {
		a[i] = make([]uint8, dx)
	}
	for i := range a {
		for j := range a[i] {
			a[i][j] = uint8(i * j)
		}
	}
	return
}

func main() {
	// なんか変な模様ができる
	pic.Show(Pic)
}
