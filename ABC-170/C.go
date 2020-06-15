package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 10000000
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}
func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}
func nextInt2() (int, int) {
	return nextInt(), nextInt()
}
func nextInt3() (int, int, int) {
	return nextInt(), nextInt(), nextInt()
}
func nextInt4() (int, int, int, int) {
	return nextInt(), nextInt(), nextInt(), nextInt()
}
func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}
func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}
func nextFloat64s(n int) []float64 {
	slice := make([]float64, n)
	for i := 0; i < n; i++ {
		slice[i] = nextFloat64()
	}
	return slice
}
func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}
func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	x, n := nextInt(), nextInt()
	p := map[int]bool{}
	for i := 0; i < n; i++ {
		pi := nextInt()
		p[pi] = true
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
	puts(ans)
}
