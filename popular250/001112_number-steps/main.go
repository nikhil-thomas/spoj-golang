// Title             : 'Number Steps'
// SPOJ ID           : '1112'
// Description       : 'https://www.spoj.com/problems/NSTEPS/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 11, 2018'
// Last Modified On  : 'August 11, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Camilo Andrés Varela León
// Date         : 2006-11-25
// Time limit   : 1.159s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: ERL JS-RHINO NODEJS PERL6 VB.NET
// Resource     : Asia - Tehran 2000
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"strconv"
)

type coords struct {
	x int
	y int
}

func newCoords(x, y int) *coords {
	return &coords{x, y}
}

func (c *coords) isCoordsHasNumber() bool {
	if c.x < 0 || c.y < 0 {
		return false
	}

	if c.x == 0 || c.x == 1 {
		return c.y == c.x
	}

	return c.y == c.x || c.y == c.x-2
}

func (c *coords) computePositionVal() string {
	if !c.isCoordsHasNumber() {
		return "No Number"
	}
	if c.x == 0 || c.x == 1 {
		return strconv.Itoa(c.x)
	}
	var val int

	val = c.x + c.y
	if c.x%2 != 0 {
		val--
	}

	return strconv.Itoa(val)
}

func main() {
	var testCount int
	fmt.Scanf("%d", &testCount)

	tests := []*coords{}
	for i := 0; i < testCount; i++ {
		testCoords := newCoords(0, 0)
		fmt.Scanf("%d %d", &testCoords.x, &testCoords.y)
		tests = append(tests, testCoords)
	}
	for _, c := range tests {
		fmt.Println(c.computePositionVal())
	}
}
