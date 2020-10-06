package main

import "fmt"

// func main() {
// 	s := make([]int, 10)
// 	fmt.Println(s)
// }

// func main() {
// 	var a [10]int
// 	fmt.Println(a)
// }

// func main() {
// 	s := make([]int, 8)
// 	len(s)
// }

func pow(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func main() {
	a := []int{1, 2, 3}
	pow(a)
	fmt.Println(a)
}
