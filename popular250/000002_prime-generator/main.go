// Title             : 'PRIME1 - Prime Generator'
// SPOJ ID           : '2'
// Description       : 'https://www.spoj.com/problems/PRIME1/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 08, 2018'
// Last Modified On  : 'August 08, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"math"
)

type input struct {
	start uint32
	end   uint32
}

var primeLookUp = make(map[uint32]bool)

func main() {
	var testCount uint32
	var inputs []input

	fmt.Scanf("%d", &testCount)

	for i := uint32(0); i < testCount; i++ {
		var in input
		fmt.Scanf("%d %d", &in.start, &in.end)
		inputs = append(inputs, in)
	}

	fmt.Printf("\n")
	for _, in := range inputs {
		printPrimes(in.start, in.end)
		fmt.Printf("\n")
	}
}

func printPrimes(start, end uint32) {
	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func isPrime(n uint32) bool {
	check, ok := primeLookUp[n]
	if ok {
		return check
	}
	if n == 1 {
		primeLookUp[n] = false
		return false
	}
	for i := uint32(2); i <= uint32(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			primeLookUp[n] = false
			return false
		}
	}
	primeLookUp[n] = true
	return true
}
