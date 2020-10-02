// package main

// import "fmt"

// // func doSomething() (a int) {
// // 	return
// // }

// // func doSomething() (x, y int) {
// // 	y = 5
// // 	return
// // }

// func doSomething() (int, int) {
// 	var x, y int
// 	y = 5
// 	return x, y
// }

// func main() {
// 	fmt.Println(doSomething())
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Printf("%#v\n", func(x, y int) int { return x + y })
// 	fmt.Printf("%#v\n", func(x, y int) int { return x + y }(2, 3))
// }

// package main

// import "fmt"

// func returnFunc() func() {
// 	return func() {
// 		fmt.Println("I'm a function")
// 	}
// }

// func main() {
// 	f := returnFunc()
// 	f()
// }

// package main

// import "fmt"

// func callFunction(f func()) {
// 	f()
// }

// func main() {
// 	callFunction(func() {
// 		fmt.Println("I'm a function")
// 	})
// }

package main

import "fmt"

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := later()

	fmt.Println(f("Golang"))
	fmt.Println(f("is"))
	fmt.Println(f("awesome!"))
}
