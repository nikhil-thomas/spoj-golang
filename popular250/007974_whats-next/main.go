// Title             : 'What’s Next'
// SPOJ ID           : '7974'
// Description       : 'spoj_problem_link'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 12, 2018'
// Last Modified On  : 'August 12, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Omar ElAzazy
// Date         : 2010-11-30
// Time limit   : 1.799s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: ASM64
// Resource     : ACPC 2010
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
)

func isArithematicProgression(v1, v2, v3 int) (bool, int) {
	if v2-v1 != v3-v2 {
		return false, 0
	}
	commonDifference := v2 - v1
	return true, commonDifference
}

func isGeometricProgression(v1, v2, v3 int) (bool, int) {
	if float32(v2)/float32(v1) != float32(v3)/float32(v2) {
		return false, 0
	}
	commonRatio := v2 / v1
	return true, commonRatio
}

type test struct {
	v1, v2, v3 int
}

func main() {
	testCases := []*test{}

	for {
		testCase := new(test)
		fmt.Scanf("%d %d %d", &testCase.v1, &testCase.v2, &testCase.v3)
		if testCase.v1 == 0 && testCase.v2 == 0 && testCase.v3 == 0 {
			break
		}
		testCases = append(testCases, testCase)
	}
	for _, testCase := range testCases {
		if ok, cd := isArithematicProgression(testCase.v1, testCase.v2, testCase.v3); ok {
			fmt.Println("AP", testCase.v3+cd)
		} else if ok, cr := isGeometricProgression(testCase.v1, testCase.v2, testCase.v3); ok {
			fmt.Println("GP", testCase.v3*cr)
		}
	}
}
