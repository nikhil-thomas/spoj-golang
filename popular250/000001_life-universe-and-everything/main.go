// Title             : 'TEST - Life, the Universe, and Everything'
// SPOJ ID           : '1'
// Description       : 'https://www.spoj.com/problems/TEST/'
// Author            : 'Nikhil Thomas'
// Created On        : 'CURRENT_MONTH_NAME 07, 2018'
// Last Modified On  : 'CURRENT_MONTH_NAME 08, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

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
