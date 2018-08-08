// Title             : 'ADDREV - Adding Reversed Numbers'
// SPOJ ID           : '42'
// Description       : 'https://www.spoj.com/problems/ADDREV/'
// Author            : 'Nikhil Thomas'
// Created On        : 'August 06, 2018'
// Last Modified On  : 'August 08, 2018'
//----------::::::::::----------::::::::::----------::::::::::----------//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var testNum int
	fmt.Scanf("%d", &testNum)

	var results []int
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < testNum; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		parts := strings.Split(text, " ")
		if len(parts) != 2 {
			continue
		}

		v1, _ := strconv.Atoi(parts[0])
		v2, _ := strconv.Atoi(parts[1])
		sum := reverseInt(v1) + reverseInt(v2)
		sum = reverseInt(sum)

		results = append(results, sum)
	}

	for _, v := range results {
		fmt.Println(v)
	}
}

func reverseInt(v int) int {
	var rev int
	for v > 0 {
		digit := v % 10
		rev = rev*10 + digit
		v /= 10
	}
	return rev
}
