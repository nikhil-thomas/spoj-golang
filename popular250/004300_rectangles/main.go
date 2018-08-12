// Title             : 'Rectangles'
// SPOJ ID           : '4300'
// Description       : 'https://www.spoj.com/problems/AE00/'
// Author            : 'Nikhil Thomas'
// Created On        : 'AUGUST 12, 2018'
// Last Modified On  : 'AUGUST 12, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Race with time
// Date         : 2009-05-03
// Time limit   : 1s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: ERL JS-RHINO NODEJS PERL6 VB.NET
// Resource     : Algorithmic Engagements 2009
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"math"
)

func countRectangles(n int) int {
	// all one row rectangles 1x1, 1x2, ... 1xn
	rCount := n

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		for j := i; j <= n/2; j++ {
			if i*j <= n {
				rCount++
			}
		}
	}
	return rCount
}

func main() {
	var nSquares int
	fmt.Scanf("%d", &nSquares)
	fmt.Println(countRectangles(nSquares))
}
