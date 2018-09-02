// Title             : 'Aggresive Cows'
// SPOJ ID           : '297'
// Description       : 'https://www.spoj.com/problems/AGGRCOW/'
// Author            : 'Nikhil Thomas'
// Created On        : 'September 02, 2018'
// Last Modified On  : 'September 02, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Roman Sol
// Date         : 2005-02-16
// Time limit   : 2s
// Source limit : 10000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All
// Resource     : USACO February 2005 Gold Division
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"sort"
)

func minDistance(positions []int, cows, start, end int) int {
	if cows == 2 {
		return positions[end] - positions[start]
	}
	mid := start + (end-start)/2
	fmt.Println(start, mid, end)
	// put one cow in mid
	// then find min distance between start and mid or mid and end
	cows--

	cowsLeftside := cows / 2
	cowsRightSide := cows / 2
	if cows%2 != 0 {
		cowsRightSide++
	}

	d1 := minDistance(positions, cowsLeftside, start, mid-1)
	d2 := minDistance(positions, cowsLeftside, mid+1, end)
	if d1 < d2 {
		return d1
	}
	return d2

}

func main() {
	cows := 3
	//stalls := 5
	x := []int{1, 2, 8, 4, 5}
	x = sort.IntSlice(x)
	fmt.Println(minDistance(x, cows, 0, len(x)-1))

}
