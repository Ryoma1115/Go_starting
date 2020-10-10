package main

import "fmt"

// import "fmt"

// type T struct {
// 	N uint
// 	_ int16
// 	S []string
// }

// func main() {
// 	t := T{
// 		N: 12,
// 		S: []string{"A", "B", "C"},
// 	}
// 	fmt.Println(t)
// }

type Point struct {
	X, Y int
}

func showStruct(s struct{ X, Y int }) {
	fmt.Println(s)
}

func main() {
	p := Point{X: 3, Y: 8}
	showStruct(p)
}
