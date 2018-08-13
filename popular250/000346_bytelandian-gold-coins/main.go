// Title             : 'Bytelandian gold coins'
// SPOJ ID           : '346'
// Description       : 'https://www.spoj.com/problems/COINS/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 13, 2018'
// Last Modified On  : 'August 13, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Tomek Czajka
// Date         : 2005-05-03
// Time limit   : 9s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: NODEJS PERL6 VB.NET
// Resource     : Purdue Programming Contest Training
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sumMap = make(map[uint32]uint32)

// maxSum computes maximum possible sum value (dollar amount) for a given bytelandian coin
// one coin 'n' can be split into 'n/2', 'n/3' and 'n/4'
// hence max dollar amount is 'n' if n >= n/2 + n/3 + n/4 else 'n/2 + n/3 + n/4'
// additional info: each of n/2, n/3, n/4 can be further divied into */2, */3 and */4
func maxSum(n uint32) uint32 {

	sum, ok := sumMap[n]

	if ok {
		return sum
	}

	if n == 0 {
		sumMap[n] = 0
		return 0
	}
	sum = maxSum(n/2) + maxSum(n/3) + maxSum(n/4)
	if sum <= n {
		sum = n
	}
	sumMap[n] = sum
	return sum

}

func main() {
	inputs := []uint32{}
	scanner := bufio.NewScanner(os.Stdin)

	// read inputs until an empty input
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			break
		}
		input, _ := strconv.ParseInt(str, 10, 32)
		inputs = append(inputs, uint32(input))
	}
	for _, input := range inputs {
		fmt.Println(maxSum(uint32(input)))
	}
}
