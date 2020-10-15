// func main() {
// 	defer func() {
// 		fmt.Println("defer")
// 	}()

// 	os.Exit(0)
// }

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("length=%d\n", len(os.Args))
	for _, v := range os.Args {
		fmt.Println(v)
	}
}
