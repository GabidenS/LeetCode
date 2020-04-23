package main

import "fmt"

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
}
func decompressRLElist(nums []int) []int {
	var res, temp []int
	for n := 0; n < len(nums); n += 2 {
		for i := 0; i < nums[n]; i++ {
			temp = append(temp, nums[n+1])
		}
		res = append(res, temp...)
		temp = append(temp[:0])

	}
	return res
}
