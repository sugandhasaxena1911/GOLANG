package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("The factorial of number is ", factorial(6))
}

func factorial(n int) int {
	var result int
	if n <= 1 {
		return 1
	}

	result = n * factorial(n-1)
	return result

}

func FirstDigit(num int) int {

	str := strconv.Itoa(num)
	n, _ := strconv.Atoi(str[:1])
	return n
}
func GenericRoot(n int) int {

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

func genericRootV2(n int) int {

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

func Reverse(num int) int {
	n := num
	reverse := 0

	for n > 0 {
		remainder := n % 10
		reverse = reverse*10 + remainder
		n = n / 10
	}
	return reverse
}
