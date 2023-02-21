package main

import "fmt"

func main() {
	str := "Hello"
	for i, v := range str {
		fmt.Println(i, v, str[i], string(v))
		fmt.Printf("%c %d \n", str[i], str[i])
	}
	//%c	the character represented by the corresponding Unicode code point
	//%d	base 10
}
