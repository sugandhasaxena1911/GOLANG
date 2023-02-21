package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 2376
	fmt.Println("The first digit is ", FirstDig(n), FirstDigv2(n))

}
func FirstDig(num int) int {

	str := strconv.Itoa(num)
	n, _ := strconv.Atoi(str[:1])
	return n
}
func FirstDigv2(num int) int {
	n := num
	for n >= 10 {
		n = n / 10
	}
	return n
}
