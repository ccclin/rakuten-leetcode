package main

import "fmt"

func Add(a, b int) (c int) {
	for b != 0 {
		c = a ^ b
		b = (a & b) << 1
		a = c
	}
	c = a
	return
}

func main() {
	a := 11
	b := 1
	ans := Add(a, b)
	fmt.Printf("%d + %d = %d", a, b, ans)
}
