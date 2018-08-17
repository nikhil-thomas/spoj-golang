// Title             : 'The Next Palindrome'
// SPOJ ID           : '5'
// Description       : 'spoj_problem_link'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 17, 2018'
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
	"strings"
)

type input struct {
	n   []rune
	mid int64
	len int64
}

func newInput(s string) *input {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, " ", "", -1)

	lenSlice := int64(len(s))
	mid := lenSlice / 2

	inp := input{
		n:   []rune(s),
		mid: mid,
		len: lenSlice,
	}

	return &inp
}

func (inp *input) String() string {
	return string(inp.n)
}

func (inp *input) isAll9s() bool {
	for _, d := range inp.n {
		if d != '9' {
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
	// if all 9s handle it as a special case
	if inp.isAll9s() {
		return inp.convert9sToP()
	}

	i := inp.mid - 1
	j := inp.mid

	if inp.len%2 == 1 {
		j = j + 1
	}

	//compare from mid and check from where the diffrence starts
	for i >= 0 && j < inp.len {
		if inp.n[i] != inp.n[j] {
			break
		}
		i--
		j++
	}

	// if the number is not already a palindrome
	// or start of difference in left greater than start of difference in right
	// n[i] > n[j]
	// copy left to right
	// 54321 -> 54345 (54 3 45)
	// 5432  -> 5445  (54 45)
	if i >= 0 && (inp.n[i] > inp.n[j]) {
		for i >= 0 {
			inp.n[j] = inp.n[i]
			i--
			j++
		}
	} else {
		// if the number is already a palidrome
		// or start io difference in left less than than start of difference in right
		// n[i] < n[j]
		// then increment n[mid] by 1
		// ripple carry to left
		// copy left to right
		// 1223221 -> 1224221 (122 4 221)
		// 123321  -> 124421  (12 33 21)
		// 12345   -> 12421   (12 4 21)
		// 123456  -> 124421  (12 44 21)
		carry := 1

		i = inp.mid - 1
		if inp.len%2 == 1 {
			inp.n[inp.mid] += rune(carry)
			carry = int((inp.n[inp.mid] - 48) / 10)
			inp.n[inp.mid] = (inp.n[inp.mid]-48)%10 + 48
			j = inp.mid + 1
		} else {
			j = inp.mid
		}
		for i >= 0 {
			inp.n[i] += rune(carry)
			carry = int((inp.n[i] - 48) / 10)
			inp.n[i] = (inp.n[i]-48)%10 + 48
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
