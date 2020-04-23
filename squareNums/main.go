package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(5))
}

func judgeSquareSum(c int) bool {
	sqt := int(math.Sqrt(float64(c)))
	if c-sqt*sqt == 0 {
		return true
	}
	rng := sqt + 1
	for i := 1; i <= rng; i++ {
		sqt2 := int(math.Sqrt(float64(c - i*i)))
		if c-i*i == sqt2*sqt2 {
			return true
		}
	}
	return false

}
