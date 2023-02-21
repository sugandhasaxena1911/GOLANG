package main

import "fmt"

func main() {

	fmt.Println("The number of digits are ", NoOfDigits(132))
}

func NoOfDigits(num int) int {
	n := num
	dg := 0
	for {
		n = n / 10
		dg = dg + 1
		if n <= 0 {
			break
		}

	}

	return dg
}
