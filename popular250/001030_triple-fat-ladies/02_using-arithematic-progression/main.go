// Title             : 'Triple Fat Ladies'
// SPOJ ID           : '1030'
// Description       : 'https://www.spoj.com/problems/EIGHTS/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 14, 2018'
// Last Modified On  : 'August 14, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Matthew Reeder
// Date         : 2006-10-30
// Time limit   : 1.197s
// Source limit : 30000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: ERL JS-RHINO NODEJS PERL6 VB.NET
// Resource     : Al-Khawarizm 2006
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"math/big"
)

func tripple8AP(n *big.Int) string {
	if n.Cmp(big.NewInt(1)) == 0 {
		return "192"
	}
	t1 := big.NewInt(int64(192))
	// commonDifference = 250
	cd := big.NewInt(int64(250))

	//tn = a + (n-1)*d
	// n-1
	n.Sub(n, big.NewInt(1))

	// (n-1)d
	n.Mul(n, cd)

	// t1 + (n-1)*d
	tn := t1.Add(t1, n)
	return tn.String()
}

func main() {
	testCount := big.NewInt(0)
	inputs := []*big.Int{}
	str := ""
	fmt.Scanf("%s", &str)
	testCount.SetString(str, 10)

	for i := big.NewInt(0); i.Cmp(testCount) != 0; i.Add(i, big.NewInt(1)) {
		str := ""
		fmt.Scanf("%s", &str)
		val := big.NewInt(0)
		val.SetString(str, 10)
		inputs = append(inputs, val)

	}

	for _, input := range inputs {
		fmt.Println(tripple8AP(input))
	}
}
