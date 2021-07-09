package main

import (
	"fmt"
	"strings"
)

func WordExists(a [][]string, b string) (ans bool) {
	ans = true
	indexA := make(map[string]int)

	for _, aIitems := range a {
		for _, item := range aIitems {
			indexA[item] += 1
		}
	}
	bSlice := strings.Split(b, "")
	for _, bItem := range bSlice {
		indexA[bItem] -= 1
		if indexA[bItem] < 0 {
			ans = false
			break
		}
	}
	return
}

func main() {
	a := [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}
	b := []string{"ABCCED", "SEE", "ABCB"}

	for _, bItem := range b {
		ans := WordExists(a, bItem)
		fmt.Printf("%t\n", ans)
	}
}
