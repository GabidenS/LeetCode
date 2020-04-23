package main

func main() {

}

func checkIfExist(arr []int) bool {
	var nums = make(map[int]int)
	for _, n := range arr {
		if n != 0 {
			nums[n] = 1
		} else if n == 0 {
			nums[n]++
		}
	}
	for _, n := range arr {
		if n == 0 && nums[n] > 1 {
			return true
		} else if n == 0 && nums[n] == 1 {
			continue
		} else {
			if _, ok := nums[n*2]; ok {
				return true
			}
		}
	}
	return false
}
