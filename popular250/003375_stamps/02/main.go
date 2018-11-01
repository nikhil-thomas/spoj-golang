package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Title             : 'Stamps'
// SPOJ ID           : '3375'
// Description       : 'https://www.spoj.com/problems/STAMPS/'
// Author            : 'Nikhil Thomas'
// Created On        : 'November 01, 2018'
// Last Modified On  : 'November 01, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Daniel Gómez Didier
// Date         : 2008-11-17
// Time limit   : 0.635s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: ERL JS-RHINO
// Resource     : 2007 - Circuito de Maratones ACIS / REDIS
//----------::::::::::----------::::::::::----------::::::::::----------//

// solution using selection sort
// wrong answer needs more work
//----------::::::::::----------::::::::::----------::::::::::----------//

type test struct {
	total int64
	fNum  int64
	set   []int64
}

func (t *test) compute() {
	fCount := 0
	for i := int64(0); i < t.fNum-i; i++ {
		maxIndex := i
		for j := i + 1; j < t.fNum-i; j++ {
			if t.set[j] > t.set[maxIndex] {
				maxIndex = j
			}
		}
		t.set[i], t.set[maxIndex] = t.set[maxIndex], t.set[i]
		fCount++
		t.total -= t.set[i]
		fmt.Println(t.set[maxIndex], t.total)
		if t.total < 0 {
			break
		}
	}
	if t.total > 0 {
		fmt.Println("impossible")
		return
	}
	fmt.Println(fCount)
}

func main() {
	testCount := int64(0)
	tests := []test{}
	s := ""
	fmt.Scanf("%s", &s)
	testCount, _ = strconv.ParseInt(s, 10, 64)

	for i := int64(0); i < testCount; i++ {
		t := test{}
		s1 := ""
		s2 := ""
		fmt.Scanf("%s %s", &s1, &s2)
		t.total, _ = strconv.ParseInt(s1, 10, 64)
		t.fNum, _ = strconv.ParseInt(s2, 10, 64)
		reader := bufio.NewReader(os.Stdin)
		line, _, _ := reader.ReadLine()
		text := strings.TrimSpace(string(line))
		vals := strings.Split(text, " ")
		for k := int64(0); k < t.fNum; k++ {
			n, err := strconv.ParseInt(vals[k], 10, 64)
			if err != nil {
				continue
			}
			t.set = append(t.set, n)
		}
		tests = append(tests, t)
	}

	for i, t := range tests {
		fmt.Print("Scenario #", i+1, ":\n")
		t.compute()
	}
}
