package main

import (
	"fmt"
	"unicode"
)

/*Write a program that accepts a sentence and calculate the number of letters and digits.

Suppose the following input is supplied to the program:

hello world! 123
Then, the output should be:

LETTERS 10
DIGITS 3*/

func main() {
	sumDigLetter("hello world! 123")
}

func sumDigLetter(input string) {
	var letter, digit int

	for v, char := range input {
		fmt.Printf("%T %v\n", v, v)       // index
		fmt.Printf("%T %v\n", char, char) // index

		fmt.Println(char, string(char))
		/*
			97-122  : a-z
			65-90   : A-Z
			0-9     : 48-57
		*/

		if unicode.IsLetter(char) {
			letter++
		}
		if unicode.IsDigit(char) {
			digit++
		}
	}

	fmt.Println(letter, digit)

}
