// Title             : 'Hangover'
// SPOJ ID           : '920'
// Description       : 'https://www.spoj.com/problems/HANGOVER/'
// Author            : 'Nikhil Thomas'
// Created On        : 'AUGUST 25, 2018'
// Last Modified On  : 'AUGUST 25, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Wanderley Guimarăes
// Date         : 2006-06-09
// Time limit   : 0.213s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: ERL JS-RHINO NODEJS PERL6 VB.NET
// Resource     : ACM Mid Central Regionals 2001
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"math"
)

var sumMap = make(map[int]float32)

func numCards(sum float32) int {

	for i := 1; i < math.MaxInt16; i++ {
		val := sumTillCardN(i)
		if val >= sum {
			return i
		}

	}
	return -1
}

func sumTillCardN(n int) float32 {
	sum, ok := sumMap[n]
	if ok {
		return sum
	}
	if n == 0 {
		return 0
	}

	sum = (1 / (float32(n) + 1)) + sumTillCardN(n-1)
	sumMap[n] = sum
	return sum
}

func main() {
	var inputs []float32
	var input float32
	for {
		fmt.Scanf("%f", &input)
		if input == 0 {
			break
		}
		inputs = append(inputs, input)
	}
	for _, input := range inputs {
		fmt.Println(numCards(input), "card(s)")
	}
}
