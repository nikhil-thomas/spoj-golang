// Title             : 'Feynman'
// SPOJ ID           : 'SAMER08F'
// Description       : 'https://www.spoj.com/problems/SAMER08F/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 11, 2018'
// Last Modified On  : 'August 11, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Diego Satoba
// Date         : 2008-11-23
// Time limit   : 0.165s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : ASM64 C CPP C++ 4.3.2 FORTRAN JAVA PAS-GPC PAS-FPC
// Resource     : South American Regional Contests 2008
//----------::::::::::----------::::::::::----------::::::::::----------//

// August 11 2018 : spoj will accept golang solution for this problem
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
)

// computeSquareCount computes nunmber of suqares in this problem
// num of squares = num of 1 sided squares + num of 2 sided squares + ... + num of n sided squares
// num of squares = n^2 + (n-1)^2 + ... + 1^2
func computeSquareCount(n int) int {
	var squares int
	for i := n; i >= 0; i-- {
		squares += i * i
	}
	return squares
}

func main() {
	tests := []*int{}
	for {
		var test int
		fmt.Scanf("%d", &test)
		if test == 0 {
			break
		}
		tests = append(tests, &test)
	}
	for _, v := range tests {
		fmt.Println(computeSquareCount(*v))
	}
}
