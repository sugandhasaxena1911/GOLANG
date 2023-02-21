package main

import "fmt"

func main() {
	//fmt.Println("The generic root is ", genericRoot(238))
	fmt.Println("the generic root is ", genericRootv2(67))
}

func genericRoot(n int) int {

	num := n
	total := 0
	for num > 0 {

		total = total + num%10
		num = num / 10
	}
	if total < 10 {
		return total
	} else {
		return genericRoot(total)

	}
}

func genericRootv2(n int) int {

	for {
		total := 0

		for num := n; num > 0; num = num / 10 {
			fmt.Println("nums", n, num)
			total = total + (num % 10)
		}
		fmt.Println("Total", total)
		if total < 10 {
			return total
		} else {
			n = total
		}
	}
}
