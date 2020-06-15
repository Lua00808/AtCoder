package main

// 引数の絶対値を返す
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
