package main

import "fmt"

func main() {

	fmt.Println("hello")
	n := 2442
	nRev := reverse(n)
	fmt.Println("the reverse of", n, "is ", nRev)
	if n == nRev {
		fmt.Println("the number is pallindrome")
	}
}
func reverse(num int) int {
	n := num
	reverse := 0

	for n > 0 {
		remainder := n % 10
		reverse = reverse*10 + remainder
		n = n / 10
	}
	return reverse
}
