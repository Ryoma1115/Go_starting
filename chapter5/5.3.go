package main

import "fmt"

type T struct {
	N uint
	_ int16
	S []string
}

func main() {
	t := T{
		N: 12,
		S: []string{"A", "B", "C"},
	}
	fmt.Println(t)
}
