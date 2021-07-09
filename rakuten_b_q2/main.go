package main

import (
	"fmt"
	"math"
	"strings"
)

const DELTA = 0.000000001
const DIFF = 1000000
const ROUND_ZERO = "000000"
const INITIAL_Z = 100.0

func Sqrt(x float64) (z float64) {
	z = INITIAL_Z

	step := func() float64 {
		return z - (z*z-x)/(2*z)
	}

	for zz := step(); math.Abs(zz-z) > DELTA; {
		z = zz
		zz = step()
	}
	return
}

func IsRootINT(x int) (ans bool) {
	rootX := Sqrt(float64(x))
	roundRootX := math.Round(rootX*DIFF) / DIFF
	roundRootXArray := strings.Split(fmt.Sprintf("%f", roundRootX), ".")
	ans = roundRootXArray[1] == ROUND_ZERO
	return
}

func main() {
	for i := 0; i < 1000; i++ {
		ans := IsRootINT(i)
		if ans {
			fmt.Printf("%d: %t\n", i, ans)
		}
	}
}
