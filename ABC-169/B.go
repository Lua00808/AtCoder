package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}

	for _, v := range num {
		if v == 0 {
			fmt.Println(0)
			return
		}
	}
	res := 1
	for _, v := range num {
		if v > 1000000000000000000/res {
			fmt.Println(-1)
			return
		} else {
			res *= v
		}
	}
	fmt.Println(res)
}
