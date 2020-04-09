package main

import "strings"

func main() {

}
func backspaceCompare(S string, T string) bool {
	var sw []string
	var tw []string
	for _, sn := range S {
		if sn == '#' {
			if len(sw) != 0 {
				sw = sw[:len(sw)-1]
			}
			continue
		}
		sw = append(sw, string(sn))
	}
	for _, tn := range T {
		if tn == '#' {
			if len(tw) != 0 {
				tw = tw[:len(tw)-1]
			}
			continue
		}
		tw = append(tw, string(tn))
	}
	strings.Join(tw, "")
	if strings.Join(sw, "") == strings.Join(tw, "") {
		return true
	}
	return false
}
