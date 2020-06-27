package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	a2 := a * a
	a3 := a * a * a
	fmt.Println(a + a2 + a3)
}
