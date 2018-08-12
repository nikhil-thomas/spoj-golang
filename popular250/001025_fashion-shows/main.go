// Title             : 'Fashion Shows'
// SPOJ ID           : '1025'
// Description       : 'https://www.spoj.com/problems/FASHION/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 12, 2018'
// Last Modified On  : 'August 12, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Matthew Reeder
// Date         : 2006-10-29
// Time limit   : 0.970s
// Source limit : 30000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: ERL JS-RHINO NODEJS PERL6 VB.NET
// Resource     : Al-Khawarizm 2006
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type test struct {
	count  int
	male   []int
	female []int
}

func (tst *test) computeSumOfHotness() int {
	sort.Ints(tst.male)
	sort.Ints(tst.female)

	var hotnessSum int
	for i, v := range tst.male {
		hotnessSum += v * tst.female[i]
	}
	return hotnessSum
}

func main() {
	testCount := 0
	testCases := []*test{}
	fmt.Scanf("%d", &testCount)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < testCount; i++ {

		testCase := new(test)
		scanner.Scan()
		text := scanner.Text()
		text = strings.Replace(text, "\n", "", -1)
		parts := strings.Split(text, " ")
		testCase.count, _ = strconv.Atoi(parts[0])

		scanner.Scan()
		text = scanner.Text()
		text = strings.Replace(text, "\n", "", -1)
		parts = strings.Split(text, " ")
		for _, v := range parts {
			htLevel, _ := strconv.Atoi(v)
			testCase.male = append(testCase.male, htLevel)
		}

		scanner.Scan()
		text = scanner.Text()
		text = strings.Replace(text, "\n", "", -1)
		parts = strings.Split(text, " ")
		for _, v := range parts {
			htLevel, _ := strconv.Atoi(v)
			testCase.female = append(testCase.female, htLevel)
		}

		testCases = append(testCases, testCase)
	}

	for _, testCase := range testCases {
		fmt.Println(testCase.computeSumOfHotness())
	}
}
