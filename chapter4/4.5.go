// package main

// import "fmt"

// func receiver(ch <-chan int) {
// 	for {
// 		i := <-ch
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	ch := make(chan int)

// 	go receiver(ch)

// 	i := 0
// 	for i < 10000 {
// 		ch <- i
// 		i++
// 	}
// }

package main

import (
	"fmt"
	"time"
)

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func main() {
	ch := make(chan int, 20)

	go receive("1st goroutine", ch)
	go receive("2st goroutine", ch)
	go receive("3st goroutine", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)

	time.Sleep(3 * time.Second)
}
