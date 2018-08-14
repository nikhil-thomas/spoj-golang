// Title             : 'Julka'
// SPOJ ID           : '54'
// Description       : 'https://www.spoj.com/problems/JULKA/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 14, 2018'
// Last Modified On  : 'August 14, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Adam Dzedzej
// Date         : 2004-06-08
// Time limit   : 2s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: NODEJS PERL6 VB.NET
// Resource     : Internet Contest Pogromcy Algorytmow (Algorithm Tamers) Round II, 2003
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"math/big"
)

type input struct {
	// use big.Int instead of big.Float to ensure
	// that the output is exact without any rounding off
	// but, handle divide by 2 explicitly
	// to append "0.5" if needed
	total, diff big.Int
}

func (ip *input) divideApples() {
	if ip.total.String() == "0" {
		fmt.Printf("%v\n%v\n", 0, 0)
		return
	}

	if ip.diff.String() == "0" {
		fmt.Printf("%v\n%v\n", ip.total.String(), 0)
		return
	}

	var klaudia big.Int
	klaudia.Add(&ip.total, &ip.diff)

	var natalia big.Int
	natalia.Sub(&ip.total, &ip.diff)

	fmt.Printf("%s\n%s\n", divideBy2(&klaudia), divideBy2(&natalia))

}

// klaudia = total/2 + diff/2 = (total + diff)/2
// natalia = total/2 - diff/2 = (total - diff)/2
func divideBy2(val *big.Int) string {
	var q, mod big.Int
	q.DivMod(val, big.NewInt(int64(2)), &mod)

	result := q.Text(10)
	if mod.Text(10) == "1" {
		result += ".5"
	}

	return result
}

func main() {
	tests := []*input{}
	for i := 0; i < 10; i++ {
		test := new(input)
		var str string
		fmt.Scanf("%s", &str)
		(test.total).SetString(str, 10)
		str = ""
		fmt.Scanf("%s", &str)
		(test.diff).SetString(str, 10)
		tests = append(tests, test)
	}
	for _, test := range tests {
		test.divideApples()
	}
}
