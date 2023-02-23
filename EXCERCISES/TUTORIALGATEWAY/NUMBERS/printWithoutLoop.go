package main

import "fmt"

func main() {

	Printnum(1, 100)

}

func Printnum(start, end int) {
	fmt.Println(start)
	if start == end {
		return
	}
	Printnum(start+1, end)

}
