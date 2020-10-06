// package main

// import "fmt"

// func init() {
// 	fmt.Println("init()")
// }

// func main() {
// 	fmt.Println("main()")
// }

package main

import "fmt"

var S = ""

func init() {
	S = S + "A"
}

func init() {
	S = S + "B"
}

func init() {
	S = S + "C"
}

func main() {
	fmt.Println(S)
}
