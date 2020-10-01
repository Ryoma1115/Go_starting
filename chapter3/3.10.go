package main

import "fmt"

func main() {
	s := "Taro" + " " + "Yamada"
	fmt.Printf("%v\n", s)
	n := 165 & 155
	fmt.Printf("%v\n", n)
	n = 92 ^ 137
	fmt.Printf("%v\n", n)
	n = 108 &^ 13
	fmt.Printf("%v\n", n)
	n = ^uint8(13)
	fmt.Printf("%v\n", n)
}
