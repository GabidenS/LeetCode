package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-123))
}
func reverse(x int) int {
	var res int
	for x != 0 {
		res += int(math.Pow10(countDigits(x)-1)) * (x % 10)
		x /= 10
	}
	if res > int(math.Pow(2, 31)-1) {
		return 0
	}
	return res
}

func countDigits(n int) int {
	i := 0
	for n != 0 {
		i++
		n /= 10
	}
	return i
}
