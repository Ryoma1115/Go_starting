// package foo

// func IsOne(n int) bool {
// 	if n == 1 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

package foo

import "testing"

func TestIsOne(t *testing.T) {
	n := 1
	b := IsOne(n)
	if b != true {
		t.Errorf("%d is not one", n)
	}
}
