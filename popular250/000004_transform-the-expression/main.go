// Title             : 'Transform the Expression'
// SPOJ ID           : '4'
// Description       : 'https://www.spoj.com/problems/ONP/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 10, 2018'
// Last Modified On  : 'August 10, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     :mima
// Date         : 2004-05-01
// Time limit   : 5s
// Source limit : 50000B
// Memory limit :	1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All except: NODEJS PERL6 VB.NET
// Info         : Two-argument operators: +, -, *, /, ^ (priority from the
//                lowest to the highest), brackets ( ).
//                Operands: only letters: a,b,...,z.
//                Assume that there is only one RPN form
//                (no expressions like a*b*c).
//----------::::::::::----------::::::::::----------::::::::::----------//

// Training Resource : Shunting Yard Algorithm
//                     https://brilliant.org/wiki/shunting-yard-algorithm/
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
	"log"
)

var precedence map[rune]int

func init() {
	precedence = make(map[rune]int)
	precedence['+'] = 1
	precedence['-'] = 2
	precedence['*'] = 3
	precedence['/'] = 4
	precedence['^'] = 5
}

type queue []rune
type stack []rune

func (q *queue) enqueue(val rune) {
	*q = append(*q, val)
}

func (q queue) String() string {
	var str string
	for _, v := range q {
		str += string(v)
	}
	return str
}

func (stk *stack) push(val rune) {
	*stk = append(stack{val}, *stk...)
}

func (stk *stack) pop() rune {
	val := (*stk)[0]
	*stk = (*stk)[1:]
	return val
}

func (stk *stack) top() rune {
	return (*stk)[0]
}

func (stk *stack) isEmpty() bool {
	return len(*stk) == 0
}

func (stk stack) String() string {
	var str string
	for _, v := range stk {
		str += string(v)
	}
	return str
}

func infixToRPN(exp string) string {
	outputQueue := queue{}
	opcodeStack := stack{}
	for _, token := range exp {
		if isOperand(token) {
			outputQueue.enqueue(token)
		} else if token == rune('(') {
			opcodeStack.push(token)
		} else if token == rune(')') {
			for !opcodeStack.isEmpty() && opcodeStack.top() != rune('(') {
				val := opcodeStack.pop()
				outputQueue.enqueue(val)
			}
			if !opcodeStack.isEmpty() && opcodeStack.top() == rune('(') {
				opcodeStack.pop()
			}
		} else if isOperator(token) {
			if !opcodeStack.isEmpty() {
				for !opcodeStack.isEmpty() && precedence[opcodeStack.top()] > precedence[token] {
					val := opcodeStack.pop()
					outputQueue.enqueue(val)
				}
			}
			opcodeStack.push(token)

		} else {
			log.Fatalln("Invalid token in input", string(token))
		}
	}
	numPops := len(opcodeStack)
	for i := 0; i < numPops; i++ {
		outputQueue.enqueue(opcodeStack.pop())
	}
	return outputQueue.String()
}

func isOperand(val rune) bool {
	if (val >= rune('a') && val <= rune('z')) || (val >= rune('A') && val <= rune('Z')) {
		return true

	}
	return false
}

func isOperator(val rune) bool {
	if val == rune('+') || val == rune('-') || val == rune('*') || val == rune('/') || val == rune('^') {
		return true
	}
	return false
}

func main() {
	var testNum int
	fmt.Scanf("%d", &testNum)

	var tests []string
	var test string
	for i := 0; i < testNum; i++ {
		fmt.Scanln(&test)
		tests = append(tests, test)
	}
	for _, infix := range tests {
		fmt.Println(infixToRPN(infix))
	}
}
