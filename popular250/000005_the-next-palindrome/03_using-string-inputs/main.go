// Title             : 'The Next Palindrome'
// SPOJ ID           : '5'
// Description       : 'spoj_problem_link'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 15, 2018'
// Last Modified On  : 'August 15, 2018'
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
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

type input struct {
	n   []int
	mid int64
	len int64
}

func newInput(s string) *input {
	nSlice := []int{}
	for _, r := range s {
		digit, _ := strconv.Atoi(string(r))
		nSlice = append(nSlice, digit)
	}

	lenSlice := int64(len(nSlice))
	mid := lenSlice / 2

	if lenSlice%2 == 0 {
		mid = mid - 1
	}

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

func (inp *input) incMidAndMirrorLeft() string {
	i := inp.mid
	carry := 1
	for i >= 0 {
		inp.n[i] += carry
		carry = inp.n[i] / 10
		inp.n[i] = inp.n[i] % 10
		i--
	}
	if carry != 0 {
		inp.n = append([]int{carry}, inp.n...)
		inp.n = append(inp.n, 0)
	}

	var j int64
	if inp.len%2 == 1 {
		i = inp.mid - 1
		j = inp.mid + 1
	} else {
		i = inp.mid
		j = inp.mid + 1
	}
	return inp.mirrorLeft(i, j)
}

func (inp *input) mirrorLeft(i, j int64) string {
	for i >= 0 && j < inp.len {
		inp.n[j] = inp.n[i]
		i--
		j++
	}

	return inp.String()
}

func (inp *input) nextPalindrome() string {
	if inp.isAll9s() {
		return inp.convert9sToP()
	}

	i := inp.mid
	j := inp.mid + 1

	if inp.len%2 == 1 {
		i = i - 1
	}
	for i >= 0 && j < inp.len {
		if inp.n[i] != inp.n[j] {
			break
		}
		i--
		j++
	}
	if i < 0 {
		return inp.incMidAndMirrorLeft()
	}

	if inp.n[i] > inp.n[j] {
		return inp.mirrorLeft(i, j)
	}

	return inp.incMidAndMirrorLeft()

}

func main() {
	str := ""
	fmt.Scanf("%s", &str)

	testCount, _ := big.NewInt(0).SetString(str, 10)

	inputs := []*input{}

	scanner := bufio.NewScanner(os.Stdin)
	for i := big.NewInt(0); i.Cmp(testCount) != 0; i.Add(i, big.NewInt(1)) {
		scanner.Scan()
		str := scanner.Text()

		inputs = append(inputs, newInput(str))
	}

	for _, input := range inputs {
		fmt.Println(input.nextPalindrome())
	}
}
