package main

func main() {

}

func twoSum(nums []int, target int) []int {
	var res []int
	temp := make(map[int]int)
	for i := 0; i < len(nums); i++ {

		temp[nums[i]] = i
		if i > 0 {
			_, ok := temp[target-nums[i]]
			if ok {
				return []int{temp[target-nums[i]], i}
			}
		}

	}
	return res
}
