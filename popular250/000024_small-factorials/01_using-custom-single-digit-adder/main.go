// Title             : FCTRL2 - Small factorials
// SPOJ ID           : 24
// Description       : 'https://www.spoj.com/problems/FCTRL2/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 08, 2018'
// Last Modified On  : 'August 09, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     :	adrian
// Date         :	2004-05-28
// Time limit   :	1s
// Source limit :	2000B
// Memory limit :	1536MB
// Cluster      :	Cube (Intel G860)
// Languages    :	All
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"strconv"
)

type bigInt []int

var factorialMap = make(map[string]bigInt)

func newBigInt(val int) bigInt {
	var bi bigInt
	for val > 0 {
		digit := val % 10
		bi = append(bigInt{digit}, bi...)
		val /= 10
	}
	return bi
}

func (bi bigInt) String() string {
	var str string
	skipZero := true
	for _, v := range bi {
		if skipZero {
			if v == 0 {
				continue
			}
			skipZero = false
		}

		str += strconv.Itoa(v)
	}
	return str
}

func addBigInts(a, b bigInt) bigInt {
	if len(a) < len(b) {
		a, b = b, a
	}
	result := make(bigInt, len(a)+1, len(a)+1)

	aIndex := len(a) - 1
	bIndex := len(b) - 1
	rIndex := len(result) - 1
	var sum, carry int
	for aIndex >= 0 {
		if bIndex >= 0 {
			sum, carry = addInt(a[aIndex], b[bIndex], result[aIndex], carry)
			bIndex--
		} else {
			sum, carry = addInt(a[aIndex], result[aIndex], carry)
		}
		result[rIndex] = sum
		aIndex--
		rIndex--
	}
	for i := aIndex; i >= 0; i-- {
		sum, carry = addInt(a[aIndex], result[aIndex], carry)
		result[rIndex] = sum
		rIndex--
	}
	return result
}

func intMultiplyBigInt(a bigInt, b int) bigInt {
	result := make(bigInt, len(a), len(a))

	for i := 1; i <= b; i++ {
		result = addBigInts(result, a)
	}

	return result
}

func addInt(digits ...int) (int, int) {
	var sum, carry int

	for _, digit := range digits {
		sum += digit
	}
	if sum >= 10 {
		carry = sum / 10
		sum %= 10
	}
	return sum, carry
}

func factorial(val int) string {
	key := strconv.Itoa(val)
	factorial, ok := factorialMap[key]

	if !ok {
		factorial = newBigInt(val)
		for i := val - 1; i >= 1; i-- {
			factorial = intMultiplyBigInt(factorial, i)
		}
		factorialMap[key] = factorial
	}
	return factorial.String()
}

func main() {
	var testNum int
	fmt.Scanf("%d", &testNum)

	var tests []int

	for i := 0; i < testNum; i++ {
		var test int
		fmt.Scanf("%d", &test)
		tests = append(tests, test)
	}
	for _, val := range tests {
		fmt.Println(factorial(val))
	}
}
