package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(arr)
	tmp := arr[:2-1]
	fmt.Println(tmp)
	arr = arr[2:]
	fmt.Println(arr)

}

// func createTargetArray(nums []int, index []int) []int {
// 	var res []int
// 	for i := 0; i < len(index); i++ {
// 		if i == 0 {
// 			res = append(res, nums[i])
// 		} else {
// 			tem := res[index[i]+1:]
// 			res = append(res[:index[i]], nums[i])
// 			res = append(res, tem...)

// 		}
// 	}
// 	return res
// }
