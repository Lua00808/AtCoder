package main

import "fmt"

func main() {
	var n int
	var div int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	loop := true
	for loop {
		for i, v := range a {
			if v%2 == 1 {
				loop = false
			}
			a[i] = a[i] / 2
		}
		if loop {
			div++
		}
	}
	fmt.Println(div)
}
