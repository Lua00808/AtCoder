package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)
	var count int
	count = 0
	for i, _ := range s {
		if s[i] != t[i] {
			count++
		}
	}
	fmt.Println(count)
}
