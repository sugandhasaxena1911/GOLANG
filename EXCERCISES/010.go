package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*Write a program which accepts a sequence of comma separated 4 digit binary numbers as its input
and then check whether they are divisible by 5 or not.
The numbers that are divisible by 5 are to be printed in a comma separated sequence.

Example: 0100,0011,1010,1001,1111

Then the output should be: 1010,1111
*/

func main() {

	s := binary_div5("1111,0100,0011,1010,1001")
	fmt.Println("output ", s)

}

func binary_div5(input string) string {
	sl := strings.Split(input, ",")
	slNew := []string{}
	for _, v := range sl {

		n, e := strconv.ParseInt(v, 2, 64)
		if e != nil {
			log.Fatal("Error ", e)
		}
		if n%5 == 0 {
			s := strconv.FormatInt(n, 2) //converted to int64 to binary
			slNew = append(slNew, s)
		}

	}

	return strings.Join(slNew, ",")
}
