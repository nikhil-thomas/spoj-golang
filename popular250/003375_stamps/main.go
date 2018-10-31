package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Title             : 'Stamps'
// SPOJ ID           : '3375'
// Description       : 'https://www.spoj.com/problems/STAMPS/'
// Author            : 'Nikhil Thomas'
// Created On        : 'October 30, 2018'
// Last Modified On  : 'October 30, 2018'
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

// Non zero execution error on submission
//----------::::::::::----------::::::::::----------::::::::::----------//

type test struct {
	total int64
	fNum  int
	set   []int
}

func (t *test) compute() {
	sort.Ints(t.set)
	fCount := 0
	for i := t.fNum - 1; i >= 0; i-- {
		fCount++
		t.total -= int64(t.set[i])
		if t.total <= 0 {
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
		s := ""
		fmt.Scanf("%s %d", &s, &t.fNum)
		t.total, _ = strconv.ParseInt(s, 10, 64)
		reader := bufio.NewReader(os.Stdin)
		line, _, _ := reader.ReadLine()
		text := strings.TrimSpace(string(line))
		vals := strings.Split(text, " ")
		for k := 0; k < t.fNum; k++ {
			n, _ := strconv.Atoi(vals[k])
			t.set = append(t.set, n)
		}
		tests = append(tests, t)
	}

	for i, t := range tests {
		fmt.Print("Scenario #", i+1, ":\n")
		t.compute()
	}
}
