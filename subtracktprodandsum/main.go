package main

import (
	"fmt"
	"math"
)

func main() {

}
func subtractProductAndSum(n int) int {
	var sum, temp int
	prod := 1
	for n != 0 {
		temp = n % 10
		prod *= temp
		sum += temp
		n /= 10
	}
	return prod - sum
	math.Sqrt
	fmt.Print
}
