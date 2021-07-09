package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) (reverseS string) {
	sArray := strings.Split(s, "")
	lenS := len(sArray)
	reverseSArray := make([]string, lenS)

	for i := 0; i < len(s); i++ {
		reverseSArray[lenS-i-1] = sArray[i]
	}

	reverseS = strings.Join(reverseSArray, "")
	return
}

func main() {
	ans := Reverse("Hello World, Hello Charlie")
	fmt.Printf("%s", ans)
}
