package main

import "fmt"

func main() {
	n := 639
	fmt.Println("the reverse of", n, "is ", reverse(n))
}

func reverse(num int) int {
	n := num
	reverse := 0

	for n > 0 {
		remainder := n % 10
		reverse = reverse*10 + remainder

	}
	return reverse
}
