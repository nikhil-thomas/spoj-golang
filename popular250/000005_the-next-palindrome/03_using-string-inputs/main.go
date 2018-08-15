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
// implementation using string inputs
// credit: https://www.geeksforgeeks.org/given-a-number-find-next-smallest-palindrome-larger-than-this-number/
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func nextPalindrome(str string) string {
	return ""
}

func main() {
	testCount := big.NewInt(0)
	str := ""
	fmt.Scanf("%s", &str)
	testCount.SetString(str, 10)

	inputs := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for i := big.NewInt(0); i.Cmp(testCount) != 0; i.Add(i, big.NewInt(int64(1))) {
		scanner.Scan()
		str = scanner.Text()
		str = strings.Replace(str, "\n", "", -1)
		inputs = append(inputs, str)
	}

	for _, input := range inputs {
		fmt.Println(nextPalindrome(input))
	}
}
