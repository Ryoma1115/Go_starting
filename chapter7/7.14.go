// // package main

// // import (
// // 	"io/ioutil"
// // 	"log"
// // 	"os"
// // )

// // func main() {
// // 	f, err := os.Open("foo.txt")
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	bs, err := ioutil.ReadAll(f)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // }

// package main

// import (
// 	"io/ioutil"
// 	"log"
// )

// func main() {
// 	bs, err := ioutil.ReadFile("foo.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := ioutil.TempFile(os.TempDir(), "foo")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString("Hello, World!\n")

	fmt.Println(f.Name())
}
