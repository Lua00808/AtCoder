package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	as := make([]int, n)
	for i := range as {
		sc.Scan()
		as[i], _ = strconv.Atoi(sc.Text())
	}
	bs := make([]int, m)
	for i := range bs {
		sc.Scan()
		bs[i], _ = strconv.Atoi(sc.Text())
	}

	sbs := make([]int, m+1)
	for i := range bs {
		sbs[i+1] += sbs[i] + bs[i]
	}

	var mi, mc int
	for i := 0; i <= n; i++ {
		ri := k - mi
		if ri < 0 {
			break
		}

		l := 0
		r := m + 1
		for r-l > 1 {
			mid := (l + r) / 2
			if sbs[mid] <= ri {
				l = mid
			} else {
				r = mid
			}
		}
		cnt := i + l
		mc = max(mc, cnt)
		if i < n {
			mi += as[i]
		}
	}
	fmt.Println(mc)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
