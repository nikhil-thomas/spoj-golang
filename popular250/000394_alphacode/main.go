// Title             : 'Alphacode'
// SPOJ ID           : '394'
// Description       : 'https://www.spoj.com/problems/ACODE/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 14, 2018'
// Last Modified On  : 'August 14, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Adrian Kuegel
// Date         : 2005-07-09
// Time limit   : 0.341s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All
// Resource     : ACM East Central North America Regional Programming Contest 2004
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var decCount = make(map[string]big.Int)

// possibleDecs solves this using dynamic programming
// possibleDecs(str) = possibleDecs(str[1]) * possibleDecs(str[1:]) + possibleDecs(str[:2]) * possibleDecs(str[2:] if str[:2] is valid letter code
// possibleDecs(str) = possibleDecs(str[1]) * possibleDecs(str[1:]) if str[:2] is not valid letter code
// use big.Innt instead of int or int64 as the input strings can have length 500
func possibleDecs(code string) *big.Int {
	if count, ok := decCount[code]; ok {
		return &count
	}
	count := big.NewInt(int64(0))
	if len(code) <= 1 {
		if code != "0" {
			decCount[code] = *big.NewInt(int64(1))
			return big.NewInt(int64(1))
		} else {
			return big.NewInt(int64(0))
		}
	}

	count = count.Mul(possibleDecs(code[0:1]), possibleDecs(code[1:]))

	subStr := code[0:2]
	if subStr[0] != '0' {
		subCode, _ := strconv.Atoi(subStr)
		if subCode > 0 && subCode <= 26 {
			count = count.Add(count, possibleDecs(code[2:]))
		}
	}
	decCount[code] = *count
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inputs := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "0" {
			break
		}
		inputs = append(inputs, text)
	}
	for _, input := range inputs {
		fmt.Println(possibleDecs(input))
	}
}
