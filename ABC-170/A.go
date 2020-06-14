package main

import (
	"fmt"
)

func main() {
	res := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&res[i])
	}
	for i, _ := range res {
		if res[i] == 0 {
			fmt.Println(i + 1)
			return
		}
	}
}
