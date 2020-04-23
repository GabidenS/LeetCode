package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMaxLength([]int{1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1}))

	// fmt.Println(findMaxLength([]int{0, 0, 1, 0, 1}))
}
func findMaxLength(nums []int) int {
	var max, count int
	m := make(map[int]int)
	m[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count--
		}
		if v, ok := m[count]; ok {
			// math.Max()
			max = int(math.Max(float64(max), float64(i-v)))
		} else {
			m[count] = i
		}

	}
	return max
}

// for i := 0; i < len(nums)/2; i++ {
// 	if nums[i*2]+nums[i*2+1] == 1 {
// 		temp++
// 	} else {
// 		temp = 0
// 	}
// 	if temp > max {
// 		max = temp
// 	}
// }
// fmt.Println("Max", max)
// temp = 0
// if len(nums) > 1 {
// 	nums = nums[1:]
// 	for i := 0; i < len(nums)/2; i++ {
// 		if nums[i*2]+nums[i*2+1] == 1 {
// 			temp++
// 		} else {
// 			temp = 0
// 		}
// 		if temp > max {
// 			max = temp
// 		}
// 	}
// }

// public class Solution {

//     public int findMaxLength(int[] nums) {
//         Map<Integer, Integer> map = new HashMap<>();
//         map.put(0, -1);
//         int maxlen = 0, count = 0;
//         for (int i = 0; i < nums.length; i++) {
//             count = count + (nums[i] == 1 ? 1 : -1);
//             if (map.containsKey(count)) {
//                 maxlen = Math.max(maxlen, i - map.get(count));
//             } else {
//                 map.put(count, i);
//             }
//         }
//         return maxlen;
//     }
// }
