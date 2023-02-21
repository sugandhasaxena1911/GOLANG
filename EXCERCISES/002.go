package main

import (
	"fmt"
	"log"
)

//which can compute the factorial of a given numbers input by user

func main() {
	var num int
	fmt.Println("Enter one number")
	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Fatal("Error errored . cannot proceed ", err)
	}
	fmt.Println("the factoral of num is ", fact(num))
}

func fact(n int) int {

	result := 1
	if n == 1 {
		return n
	}

	result = n * fact(n-1)
	return result
}
