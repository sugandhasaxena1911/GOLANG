package main

import "fmt"

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
