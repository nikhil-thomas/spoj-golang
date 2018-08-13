// Title             : 'The last digit'
// SPOJ ID           : '3442'
// Description       : 'https://www.spoj.com/problems/LASTDIG/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 13, 2018'
// Last Modified On  : 'August 13, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Jose Daniel Rodriguez Morales
// Date         : 2008-12-01
// Time limit   : 0.340s
// Source limit : 700B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: GOSU
// Resource     : Own
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
)

type input struct {
	base int
	exp  uint32
}

// last digit of power repeats after every 4th power
// 3^1 =   3, 3^5 =  24 3
// 3^2 =   9, 3^6 =  72 9
// 3^3 = 2 7, 3^7 = 218 7
// 3^4 = 8 1, 3^8 = 656 1
func (ip *input) lastDigit() int {
	if ip.exp == 0 {
		return 1
	}
	unit := ip.base % 10
	if unit == 0 || unit == 5 || unit == 1 {
		return unit
	}
	exp := int(ip.exp % 4)
	if exp == 0 {
		exp = 4
	}
	output := 1
	for i := 1; i <= exp; i++ {
		output *= unit
	}
	return output % 10
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	t := []*input{}
	for i := 0; i < n; i++ {
		test := new(input)
		fmt.Scanf("%d %d", &test.base, &test.exp)
		t = append(t, test)
	}
	for _, test := range t {
		fmt.Println(test.lastDigit())
	}
}
