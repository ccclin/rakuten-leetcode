package main

import "fmt"

func Overlapping(a [][]int, b []int) (ans [][]int) {
	for _, aItem := range a {
		if b[0] <= aItem[0] && aItem[0] <= b[1] {
			b[0] = min(aItem[0], b[0])
			b[1] = max(aItem[1], b[1])
		} else if b[0] <= aItem[1] && aItem[1] <= b[1] {
			b[0] = min(aItem[0], b[0])
			b[1] = max(aItem[1], b[1])
		} else {
			ans = append(ans, aItem)
		}
	}
	ans = append(ans, b)
	return
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

func main() {
	a := [][]int{{1, 3}, {6, 9}}
	b := []int{2, 5}

	ans1 := Overlapping(a, b)
	fmt.Printf("%+v\n", ans1)

	c := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	d := []int{4, 9}

	ans2 := Overlapping(c, d)
	fmt.Printf("%+v\n", ans2)
}
