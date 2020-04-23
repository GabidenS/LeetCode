package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	m := make(map[string]int)
	m2 := make(map[string]int)
	f, err := ioutil.ReadFile("base.txt")
	if err != nil {
		log.Fatal(err)
	}
	f2, err := ioutil.ReadFile("find.txt")
	if err != nil {
		log.Fatal(err)
	}
	finds := strings.Split(string(f2), "\n")
	strs := strings.Split(string(f), "\n")
	for i, str := range strs {
		m[str] = i
	}
	for i, str := range finds {
		m2[str] = i
	}
	fmt.Println(len(m))
	fmt.Println(finds)
	for _, fs := range finds {
		if _, ok := m[fs]; !ok {
			fmt.Println(fs, m2[fs])
		}
	}
}
