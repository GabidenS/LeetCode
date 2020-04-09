package main

func main() {

}
func maxProfit(prices []int) int {
	var profit int
	candidate := -1
	for i := 0; i < len(prices); i++ {
		if i < len(prices)-1 {
			if prices[i] < prices[i+1] && candidate == -1 {
				candidate = prices[i]
				continue
			} else if candidate != -1 {
				if prices[i] > prices[i+1] {
					profit += prices[i] - candidate
					candidate = 0
				}
			}
		} else {
			if candidate != -1 {
				profit += prices[i] - candidate
			}
		}
	}
	return profit
}

// func maxProfit(prices []int) int {
// 	var profit, buy int
// 	min := -1
// 	max := -1

// 	for i := 0; i < len(prices); i++ {
// 		if min == -1 {
// 			min = prices[i]
// 		} else {
// 			if prices[i] < min {
// 				min = prices[i]
// 			} else {
// 				buy = min
// 				max = prices[i]
// 			}
// 		}
// 		if buy != -1 {
// 			if prices[i] >= max {
// 				max = prices[i]
// 			} else {
// 				profit += max - buy
// 				buy = -1
// 				max = -1
// 				min = prices[i]
// 			}
// 		}
// 	}
// 	if buy != -1 {
// 		profit += max - buy
// 	}
// 	return profit
// }

// func maxProfit(prices []int) int {
// 	var profit int
// 	candidate := -1
