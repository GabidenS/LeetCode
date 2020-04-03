package main

func main() {

}

func maxSubArray(nums []int) int {
	var temp int
	var max int

	for i, n := range nums {
		if i == 0 {
			max = n
			temp = n
			continue
		}
		temp += n
		if temp > n {
			if temp > max {
				max = temp
			}
		} else {
			temp = n
			if temp > max {
				max = temp
			}
		}
	}

	return max
}

// func maxSubArray(nums []int) int {
// 	var temp int
// 	var max int
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}
// 	for _, n := range nums {

// 		temp += n
// 		if temp > n {
// 			if temp > max {
// 				max = temp
// 			}
// 		} else {
// 			temp = 0

// 		}

// 	}
// 	return max
// }
