package main

import "fmt"

func main() {
	arr := []int{0, 5, 5, 6, 2, 0, 0, 0, 0, 1, 4, 0}
	fmt.Println(arr)
	moveZeroes(arr)
	fmt.Println(arr)

}
func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}

// func moveZeroes(nums []int) {
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			for n := i + 1; n < len(nums); n++ {
// 				if nums[n] != 0 {
// 					nums[i], nums[n] = nums[n], nums[i]
// 					break
// 				}
// 			}
// 		}
// 	}
// }
