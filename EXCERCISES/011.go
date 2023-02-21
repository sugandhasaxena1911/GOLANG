package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*Write a program, which will find all such numbers between 100 and 300 (both included)
such that each digit of the number is an even number.
The numbers obtained should be printed in a comma-separated sequence on a single line.
*/

func main() {
	fmt.Println(eachDigEven(100, 300))
}

func eachDigEven(low, high int) string {
	sl := []string{}
	for x := low; x <= high; x++ {
		s := strconv.Itoa(x)
		fmt.Println(s)

		chk := 1

		for j := 0; j < len(s); j++ {
			sDigit := s[j : j+1]
			n, _ := strconv.Atoi(sDigit)
			if n%2 != 0 {
				chk = 0
				break
			}
		}

		if chk == 1 {
			sl = append(sl, s)
		}
	}

	fmt.Println(sl)
	return strings.Join(sl, ",")
}
