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

func tripple8BruteForce(n *big.Int) string {
	count := big.NewInt(0)
	val := big.NewInt(1)
	ans := ""
	for count.Cmp(n) != 0 {
		temp := big.NewInt(0)
		temp.SetString(val.String(), 10)
		temp.Mul(temp, val)
		temp.Mul(temp, val)
		str := temp.String()
		if len(str) >= 3 {
			str = str[len(str)-3:]
			if str == "888" {
				count.Add(count, big.NewInt(1))
				fmt.Println(val.String())
				ans = val.String()
			}
		}

		val.Add(val, big.NewInt(1))
	}
	return ans
}

func main() {
	n := big.NewInt(int64(15))
	fmt.Println(tripple8BruteForce(n))
}
