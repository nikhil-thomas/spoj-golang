// Title             : FCTRL2 - Small factorials
// SPOJ ID           : 24
// Description       : 'https://www.spoj.com/problems/FCTRL2/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 09, 2018'
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

// Training Resource : Factorial of a large number
//                     https://www.geeksforgeeks.org/factorial-large-number/
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"strconv"
)

type bigInt []int

func newBigInt(val int) bigInt {
	var bi bigInt
	for val > 0 {
		digit := val % 10
		bi = append(bi, digit)
		val /= 10
	}
	return bi
}

func (bi bigInt) String() string {
	var str string

	for i := len(bi) - 1; i >= 0; i-- {
		str += strconv.Itoa(bi[i])
	}
	return str
}

func intMultiplyBigInt(bi bigInt, n int) bigInt {
	product := bigInt{}
	var carry, productDigit int
	for _, v := range bi {
		val := v*n + carry
		productDigit = val % 10
		product = append(product, productDigit)

		carry = val / 10
	}
	if carry > 0 {
		product = append(product, newBigInt(carry)...)
	}
	return product
}

func factorial(val int) string {
	fact := bigInt{1}

	for i := 2; i <= val; i++ {
		fact = intMultiplyBigInt(fact, i)
	}
	return fact.String()
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
