package main

func main() {
	// fmt.Println(stringShift("Hello"))
}

func stringShift(s string, shift [][]int) string {

	if len(shift) > 0 {
		for _, shft := range shift {
			if shft[0] == 1 {
				s = s[len(s)-shft[1]:] + s[:len(s)-shft[1]]
			} else {
				s = s[shft[1]:] + s[:shft[1]]
			}
		}
	}

	return s
}
