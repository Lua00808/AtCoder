package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}
	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		m[num[i]]++
	}
	j := []int{}
	for _, v := range m {
		j = append(j, v)
	}
	sort.Ints(j)
	boll := 0
	for i := 0; i < (len(j) - k); i++ {
		boll += j[i]
	}
	fmt.Println(boll)
}
