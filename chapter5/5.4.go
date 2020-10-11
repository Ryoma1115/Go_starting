// package main

// import "fmt"

// type MyError struct {
// 	Message string
// 	ErrCode int
// }

// func (e *MyError) Error() string {
// 	return e.Message
// }

// func RaiseError() error {
// 	return MyError{Message: "エラーが発生しました。", ErrCode: 1234}
// }

// func main() {
// 	err := RaiseError()
// 	fmt.Println(err.Error())
// }
package main

import "fmt"

type I interface {
	Method1() string
	method2() string
}

type T struct{}

func (t *T) Method1() string {
	return "*Method()"
}

func (t *T) method2() string {
	return "*method2()"
}

func NewI() I {
	return &T{}
}

func main() {
	t := main.NewI()
	t.Method1()
	t.method2()

	fmt.Println(t)
}
