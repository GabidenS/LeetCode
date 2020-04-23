package main

import (
	"math"
	"sort"
)

func main() {

}
func lastStoneWeight(stones []int) int {

	if len(stones) == 0 {
		return 0
	}
	sort.Ints(stones)
	f

	return recursion(stones)
}

func recursion(s []int) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return s[0]
	}
	var delta int
	temp := s[:len(s)-2]
	delta = s[len(s)-1] - s[len(s)-2]

	if delta != 0 {
		if delta < 0 {
			delta *= -1
		}
		temp = append(temp, delta)
	}
	res := recursion(temp)
	return res
	math.Abs
}
