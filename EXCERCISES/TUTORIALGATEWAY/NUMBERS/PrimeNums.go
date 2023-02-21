package main

import "fmt"

func main() {
	n := 100
	fmt.Println("The prime number are ", prime(n))

}
func prime(n int) []int {
	sl := []int{}
	for x := 1; x <= n; x++ {
		check := 1
		//find if x is prime
		for y := 2; y < x; y++ {
			if x%y == 0 {
				check = 0
				break
			}
		}
		if check == 1 {
			sl = append(sl, x)
		}
	}
	return sl
}
