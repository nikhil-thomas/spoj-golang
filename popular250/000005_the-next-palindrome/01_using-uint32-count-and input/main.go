// Title             : 'The Next Palindrome'
// SPOJ ID           : '5'
// Description       : 'spoj_problem_link'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 15, 2018'
// Last Modified On  : 'August 15, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : adrian
// Date         : 2004-05-01
// Time limit   : 2s-9s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: NODEJS PERL6
//----------::::::::::----------::::::::::----------::::::::::----------//

//----------::::::::::----------::::::::::----------::::::::::----------//
// trial implementation
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func nextPalindrome(n uint32) uint32 {
	n++
	for {
		if isPalindrome(n) {
			break
		}
		n++
	}
	return n
}

func isPalindrome(n uint32) bool {
	if n == 0 {
		return true
	}

	str := ""
	for n > 0 {
		digit := int(n % 10)
		str = strconv.Itoa(digit) + str
		n = n / 10
	}

	i := 0
	j := len(str) - 1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	testCount := big.NewInt(0)
	str := ""
	fmt.Scanf("%s", &str)
	testCount.SetString(str, 10)

	inputs := []uint32{}

	for i := big.NewInt(0); i.Cmp(testCount) != 0; i.Add(i, big.NewInt(int64(1))) {
		var val uint32
		fmt.Scanf("%d", &val)
		inputs = append(inputs, val)
	}

	for _, input := range inputs {
		fmt.Println(nextPalindrome(input))
	}
}
