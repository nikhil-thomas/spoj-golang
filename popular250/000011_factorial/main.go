// Title             : FCTRL - Factorial
// SPOJ ID           : 11
// Description       : 'https://www.spoj.com/problems/FCTRL/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 08, 2018'
// Last Modified On  : 'August 08, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// Number of 0s in factorial is equal to num of 10s (2 * 5) that can be formed
// which will be equal to total number of 5s (factor 5) in the number
// hence the number of 5s (factor 5) in N will give the num of 0s in N!
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
)

func main() {
	var testNum int
	fmt.Scanf("%d", &testNum)

	var tests []uint32

	for i := 0; i < testNum; i++ {
		var test uint32
		fmt.Scanf("%d", &test)
		tests = append(tests, test)
	}

	for _, v := range tests {
		fmt.Println(zerosInFactorial(v))
	}

}

// zerosInFactorial finds number of 5s (factor 5s) in a number
func zerosInFactorial(n uint32) int {
	div := uint32(5)
	numZeroes := 0
	for n > 0 {
		numZeroes += int(n / div)
		// we need to count sub factor 5s also
		// hence divide n /5
		n = n / div
	}

	return numZeroes
}
