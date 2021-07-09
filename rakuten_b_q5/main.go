package main

import "fmt"

func Add(a, b int) (c int) {
	sliceA := make([]int, a)
	sliceB := make([]int, b)

	sliceC := append(sliceA, sliceB...)
	c = len(sliceC)
	return
}

func main() {
	a := 1
	b := 2
	ans := Add(a, b)
	fmt.Printf("%d + %d = %d", a, b, ans)
}
