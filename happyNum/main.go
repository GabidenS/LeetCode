package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := os.Args[1:]

	num, _ := strconv.Atoi(arr[0])

	fmt.Println(isHappy(num))
}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return isHappy(res)
}
