// Title             : 'The Next Palindrome'
// SPOJ ID           : '5'
// Description       : 'spoj_problem_link'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 16, 2018'
// Last Modified On  : 'August 17, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : adrian
// Date         : 2004-05-01
// Time limit   : 2s-9s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: NODEJS PERL6
//----------::::::::::----------::::::::::----------::::::::::----------//

//----------::::::::::----------::::::::::----------::::::::::----------//
// implementation using string inputs
// credit: https://www.geeksforgeeks.org/given-a-number-find-next-smallest-palindrome-larger-than-this-number/
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type input struct {
	n   []int
	mid int64
	len int64
}

func newInput(s string) *input {
	nSlice := []int{}
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, " ", "", -1)
	for _, r := range s {
		digit, _ := strconv.Atoi(string(r))
		nSlice = append(nSlice, digit)
	}

	lenSlice := int64(len(nSlice))
	mid := lenSlice / 2

	inp := input{
		n:   nSlice,
		mid: mid,
		len: lenSlice,
	}

	return &inp
}

func (inp *input) String() string {
	str := ""
	for _, d := range inp.n {
		str += strconv.Itoa(d)
	}
	return str
}

func (inp *input) isAll9s() bool {
	for _, d := range inp.n {
		if d != 9 {
			return false
		}
	}
	return true
}

func (inp *input) convert9sToP() string {
	str := ""
	for i := int64(0); i < inp.len-1; i++ {
		str += "0"
	}
	str = "1" + str + "1"
	return str
}

func (inp *input) nextPalindrome() string {
	if inp.isAll9s() {
		return inp.convert9sToP()
	}

	i := inp.mid - 1
	j := inp.mid

	if inp.len%2 == 1 {
		j = j + 1
	}
	for i >= 0 && j < inp.len {
		if inp.n[i] != inp.n[j] {
			break
		}
		i--
		j++
	}

	if i >= 0 && (inp.n[i] > inp.n[j]) {
		for i >= 0 {
			inp.n[j] = inp.n[i]
			i--
			j++
		}
	} else {
		carry := 1

		i = inp.mid - 1

		if inp.len%2 == 1 {
			inp.n[inp.mid] += carry
			carry = inp.n[inp.mid] / 10
			inp.n[inp.mid] = inp.n[inp.mid] % 10
			j = inp.mid + 1
		} else {
			j = inp.mid
		}
		for i >= 0 {
			inp.n[i] += carry
			carry = inp.n[i] / 10
			inp.n[i] = inp.n[i] % 10
			inp.n[j] = inp.n[i]
			j++
			i--
		}
	}
	return inp.String()
}

func main() {
	str := ""
	fmt.Scanln(&str)
	testCount, _ := big.NewInt(0).SetString(str, 10)

	inputs := []*input{}

	for i := big.NewInt(0); i.Cmp(testCount) != 0; i.Add(i, big.NewInt(1)) {
		fmt.Scanln(&str)
		inputs = append(inputs, newInput(str))
	}
	for _, input := range inputs {
		fmt.Println(input.nextPalindrome())
	}
}
