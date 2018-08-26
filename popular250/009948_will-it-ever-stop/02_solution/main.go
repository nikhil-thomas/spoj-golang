// Title             : 'Will it ever stop'
// SPOJ ID           : '9948'
// Description       : 'https://www.spoj.com/problems/WILLITST/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 26, 2018'
// Last Modified On  : 'August 26, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Krzysztof Lewko
// Date         : 2011-11-09
// Time limit   : 0.906s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: ASM64
// Resource     : AMPPZ 2011
//----------::::::::::----------::::::::::----------::::::::::----------//

// based on
// 01_study_the_pattern/output study_the_series.txt
//
// simplest observation: the series will end if the starting number is a poser of 2
// will it end? : yes, if the number is a poser of 2
//
// and also (the series will not end if it produces number which are multiples of 12, 6 and 3 in that order)

// to check if a numer is a power of 2
// we can divide a number by 2 continoulsly until it becomes 0 without generating any remainders in any division
// we can take log with base 2 the the result should be an integer
//
// we can do this using bit manipulation
// credit: https://www.geeksforgeeks.org/program-to-find-whether-a-no-is-power-of-two/
//
// method 1: all bits except the msb will be  0 for all powers of 2
//
// method 2: subtract 1 from n and do bitwise and of n and n-1 it result will be 0 for powers of 2

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)

	// create big integer input
	n, _ := big.NewInt(0).SetString(str, 10)

	// subtract 1 from input
	m, _ := big.NewInt(0).SetString(str, 10)
	m.Sub(m, big.NewInt(int64(1)))

	// do bitwise and of input and input-1
	result := m.And(m, n)

	// if the result of (input & (input-1)) is 0
	// input is a power of 0 and the series will end
	if result.String() == "0" {
		fmt.Println("TAK")
	} else {
		fmt.Println("NIE")
	}
}
