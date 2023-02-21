package main

import (
	"fmt"
	"log"
)

//write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included),
//and then the program should print the map with representation of the value

func main() {
	var num int
	fmt.Println("Enter one number")
	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Fatal("Error errored . cannot proceed ", err)
	}
	fmt.Println("the square of numbers  are ", square_n(num))
	//func using 'make'
	fmt.Println("the square of numbers  are ", square_nV2(num))

}

func square_n(num int) map[int]int {
	var result map[int]int
	//result := map[int]int{} // note this line , map has to be initialized . the previous usage gives error

	for x := 1; x <= num; x++ {
		result[x] = x * x
	}

	return result
}

func square_nV2(num int) map[int]int {
	var result = make(map[int]int, num) // note the equality sign and not :=  , because is used

	for x := 1; x <= num; x++ {
		result[x] = x * x
	}

	return result
}
