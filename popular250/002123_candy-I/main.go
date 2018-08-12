// Title             : 'Candy I'
// SPOJ ID           : '2123'
// Description       : 'https://www.spoj.com/problems/CANDY/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 12, 2018'
// Last Modified On  : 'August 12, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Fudan University Problem Setters
// Date         : 2007-12-01
// Time limit   : 1s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: C99 ERL JS-RHINO
// Resource     : IPSC 1999
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"sort"
)

type candybags []int

func (cb *candybags) total() int {
	var candyTotal int
	for _, bag := range *cb {
		candyTotal += bag
	}
	return candyTotal
}

func (cb *candybags) equalCount() int {
	return cb.total() / len(*cb)
}

func (cb *candybags) isEqualizable() bool {
	return cb.total()%len(*cb) == 0
}

func (cb *candybags) isEqualized() bool {
	equalTarget := cb.equalCount()
	for _, bag := range *cb {
		if bag != equalTarget {
			return false
		}
	}
	return true
}

func (cb *candybags) equqlize() int {
	if cb.isEqualized() {
		return 0
	}
	if !cb.isEqualizable() {
		return -1
	}

	var moves int

	sort.Ints(*cb)

	i, j := 0, len(*cb)-1
	equalDist := cb.equalCount()

	for i < j {
		if (*cb)[j] > equalDist && (*cb)[i] < equalDist {
			(*cb)[j]--
			(*cb)[i]++
			moves++
		}
		if (*cb)[j] == equalDist {
			j--
		}
		if (*cb)[i] == equalDist {
			i++
		}
	}

	return moves
}

func main() {
	tests := []*candybags{}

	for {
		cb := new(candybags)
		bagcount := 0
		fmt.Scanf("%d", &bagcount)
		if bagcount == -1 {
			break
		}
		for i := 0; i < bagcount; i++ {
			var candyCount int
			fmt.Scanf("%d", &candyCount)
			(*cb) = append(*cb, candyCount)
		}
		tests = append(tests, cb)
	}
	for _, test := range tests {
		fmt.Println(test.equqlize())
	}
}
