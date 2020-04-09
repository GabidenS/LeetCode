package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(countElements([]int{1, 3, 2, 3, 5, 0, 4}))

}

func countElements(arr []int) int {
	var res int
	m := make(map[int]int)

	for _, n := range arr {
		_, ok := m[n]
		if !ok {
			m[n] = 1
		} else {
			m[n]++
		}
	}
	fmt.Println(m)
	var keys []int
	for k, v := range m {
		keys = append(keys, k)
		fmt.Println(v, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(m[k], k)
		v, ok := m[k]
		_, ok1 := m[k+1]
		if ok && ok1 {
			res++
			res += v - 1
		}
	}
	return res
}
