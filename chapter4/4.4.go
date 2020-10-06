package main

import "fmt"

func main() {
	// m := make(map[int]string)

	// m[1] = "US"
	// m[81] = "Japan"
	// m[86] = "China"

	m := make(map[string]string)

	m["Yamada"] = "Taro"
	m["Osaka"] = "Hanako"
	m["Sato"] = "Jiro"

	fmt.Println(m)
}
