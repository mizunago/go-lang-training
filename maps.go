package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// Ruby でいう Hash, Python の dict 的なもの？
	var hoge map[string]int
	hoge = make(map[string]int)
	hoge["test"] = 1
	hoge["test2"] = 2
	fmt.Println(hoge)
	fmt.Println(hoge["test"])
	fmt.Println(hoge["no value"])
	fmt.Println()

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Bell Labs"].Lat)
	fmt.Println(m["Bell Labs"].Long)
}
