package main

import (
	"fmt"
	"sort"
	"strings"
)

/*Write a program that accepts a sequence of whitespace separated words as input and prints the words after removing all duplicate words and sorting them alphanumerically.

Suppose the following input is supplied to the program:

hello world and practice makes perfect and hello world again
Then, the output should be:

again and hello makes perfect practice world*/

func main() {

	fmt.Println(rem_dup_sort("hello world and practice makes perfect and hello world again"))
}
func rem_dup_sort(input string) string {
	//remove duplicate by converting to map
	//not sure for length
	//m := make([string]bool)

	m := map[string]bool{}

	sl := strings.Fields(input)
	fmt.Println("slice", sl)

	for _, v := range sl {
		fmt.Println("value is : ", v)
		m[v] = true
	}

	fmt.Println(m)
	//convert map to slice
	sl = []string{}
	for key, _ := range m {
		sl = append(sl, key)
	}
	fmt.Println("slice after removal of dups : ", sl)

	sort.Strings(sl)
	fmt.Println(sl)
	output := strings.Join(sl, " ")
	return output
}
