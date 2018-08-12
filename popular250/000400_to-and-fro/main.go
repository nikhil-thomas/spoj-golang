// Title             : 'To and Fro'
// SPOJ ID           : '400'
// Description       : 'https://www.spoj.com/problems/TOANDFRO/'
// Author            : 'Nikhil Thomas'
// Created On        : 'AUGUST 12, 2018'
// Last Modified On  : 'AUGUST 12, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

// SPOJ Specifications ---------::::::::::----------::::::::::----------//
// Added by     : Adrian Kuegel
// Date         : 2005-07-09
// Time limit   : 0.959s
// Source limit : 50000B
// Memory limit : 1536MB
// Cluster      : Cube (Intel G860)
// Languages    : All
// Resource     : ACM East Central North America Regional Programming Contest 2004
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"fmt"
)

type code struct {
	cols         int
	cipher       string
	cipherMatrix string
	message      string
}

// processInput reads the cipher and constructs the cipher matrix
// the cippher matrix is represented using a string (row major)
func (cd *code) processInput() {
	var sub string

	flag := false

	for i, letter := range cd.cipher {
		if !flag {
			sub += string(letter)
		} else {
			sub = string(letter) + sub
		}
		if i != 0 && (i+1)%cd.cols == 0 {
			flag = !flag
			cd.cipherMatrix += sub
			sub = ""
		}
	}
}

// readMessage reads constructs the message string from cipher matrix.
// readMesage reads cipherMatrix in coloumn major order
func (cd *code) readMessage() {
	var offset int

	totalRows := len(cd.cipherMatrix) / cd.cols

	for i := 0; i < len(cd.cipherMatrix); i++ {
		row := i % totalRows
		offset = i / totalRows
		readIndex := row*cd.cols + offset
		cd.message += string(cd.cipherMatrix[readIndex])

	}
}

func (cd *code) decode() string {
	cd.processInput()
	cd.readMessage()
	return cd.message
}

func main() {

	testCases := []*code{}

	for {
		var cols int
		fmt.Scanf("%d", &cols)
		if cols == 0 {
			break
		}
		test := &code{
			cols: cols,
		}
		fmt.Scanf("%s", &test.cipher)
		testCases = append(testCases, test)
	}
	for _, test := range testCases {
		fmt.Println(test.decode())
	}
}
