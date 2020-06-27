package main

// 引数の絶対値を返す
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 大きい方を返す
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 小さい方を返す
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
