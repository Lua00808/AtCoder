package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var x, n int
	fmt.Scan(&x, &n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}

	p := map[int]bool{}
	for i := 0; i < n; i++ {
		p[num[i]] = true
	}
	ans := -100
	for i := 0; i <= 101; i++ {
		if p[i] {
			continue
		}
		if abs(x-i) < abs(x-ans) {
			ans = i
		}
	}
	fmt.Println(ans)
}
