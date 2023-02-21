package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	/*Write a program that accepts a comma separated sequence of words as input
	and prints the words in a comma-separated sequence after sorting them alphabetically.
	Suppose the following input is supplied to the program:

	without,hello,bag,world
	Then, the output should be:

	bag,hello,without,world*/
	var str string
	fmt.Println("enter any comma separted string")
	fmt.Scanln(&str)
	sl := strings.Split(str, ",")
	sort.Strings(sl)
	fmt.Println(sl)
	str = strings.Join(sl, ",")
	fmt.Println(str)
	// Therefore we used split and join

}
