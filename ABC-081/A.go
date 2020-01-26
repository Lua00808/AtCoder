package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s string
	var a int
	fmt.Scan(&s)
	fmt.Println(reflect.TypeOf(s))
	if s[0] == '1' {
		a++
		fmt.Println(reflect.TypeOf(s[0]))
	}
	if s[1] == '1' {
		a++
	}
	if s[2] == '1' {
		a++
	}
	fmt.Println(a)
}
