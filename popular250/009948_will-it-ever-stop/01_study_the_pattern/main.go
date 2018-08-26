// Title             : 'Will it ever stop'
// SPOJ ID           : '9948'
// Description       : 'https://www.spoj.com/problems/WILLITST/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 26, 2018'
// Last Modified On  : 'August 26, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Krzysztof Lewko
// Date         : 2011-11-09
// Time limit   : 0.906s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: ASM64
// Resource     : AMPPZ 2011
//----------::::::::::----------::::::::::----------::::::::::----------//

// study pattern
// 01_study_the_pattern/output study_the_series.txt

package main

import (
	"fmt"
)

func mysteryCode(n int) {
	fmt.Printf("n:%d, series: ", n)
	for i := 0; i < 30; i++ {
		if n <= 1 {
			break
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 3
		}
		fmt.Printf("%d, ", n)
	}
	fmt.Printf("\n")
}

func main() {
	for i := 2; i <= 300; i++ {
		mysteryCode(i)
	}
}
