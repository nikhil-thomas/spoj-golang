package main

import (
	"fmt"
)

func main() {
	var val int
	var vals []int
	for {
		fmt.Scanln(&val)
		if val == 42 {
			break
		}
		vals = append(vals, val)
	}
	for _, v := range vals {
		fmt.Println(v)
	}
}
