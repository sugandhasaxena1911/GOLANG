package main

import "fmt"

func main() {

	printnum(1, 100)

}

func printnum(start, end int) {
	fmt.Println(start)
	if start == end {
		return
	}
	printnum(start+1, end)

}
