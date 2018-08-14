// Title             : 'Candy III'
// SPOJ ID           : '2148'
// Description       : 'https://www.spoj.com/problems/CANDY3/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 14, 2018'
// Last Modified On  : 'August 14, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Fudan University Problem Setters
// Date         : 2007-12-01
// Time limit   : 1s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: C99 ERL JS-RHINO
// Resource     : IPSC 2006
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {

	results := []string{}

	testCount := 0

	fmt.Scanf("%d", &testCount)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < testCount; i++ {
		// read empty line as per input spec
		text, _ := reader.ReadString('\n')

		sum := big.NewInt(0)
		len := big.NewInt(0)

		// read length of each test set
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		len.SetString(text, 10)

		// creation iteration variable
		// operation on len will change it
		// hence we need topreserve it and use j for iteration
		j := big.NewInt(0)
		j.SetString(len.String(), 10)

		for j.Cmp(big.NewInt(int64(0))) != 0 {
			text, _ = reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			val := big.NewInt(0)
			val.SetString(text, 10)

			sum.Add(sum, val)
			j.Sub(j, big.NewInt(int64(1)))
		}

		result := "NO"

		if len.Cmp(big.NewInt(int64(0))) != 0 {
			mod := big.NewInt(0)
			mod.Mod(sum, len)

			if mod.String() == "0" {
				result = "YES"
			}
		}
		results = append(results, result)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
