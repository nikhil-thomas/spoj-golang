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
// implementation using big.Int
// SPOJ Time Limit Exceeded
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func nextPalindrome(n *big.Int) string {
	for {
		n.Add(n, big.NewInt(int64(1)))
		if isPalindrome(n) {
			break
		}

	}
	return n.String()
}

func isPalindrome(n *big.Int) bool {
	if n.Cmp(big.NewInt(0)) == 0 {
		return true
	}

	str := n.String()

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

	inputs := []*big.Int{}

	scanner := bufio.NewScanner(os.Stdin)
	for i := big.NewInt(0); i.Cmp(testCount) != 0; i.Add(i, big.NewInt(int64(1))) {
		val := big.NewInt(0)
		scanner.Scan()
		str = scanner.Text()
		val.SetString(str, 10)
		inputs = append(inputs, val)
	}

	for _, input := range inputs {
		fmt.Println(nextPalindrome(input))
	}
}
