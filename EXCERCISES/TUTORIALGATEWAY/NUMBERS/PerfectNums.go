package main

import "fmt"

func main() {
	n := 10000
	fmt.Println("The perfect numbers till", n, "are", PerfectNums(n))
}

func PerfectNums(num int) []int {
	perfect := []int{}

	for x := 1; x <= num; x++ {
		// check if x is perfect
		sl := divisors(x)
		fmt.Println(x, sl)
		sum := 0
		for _, v := range sl {
			sum = sum + v
		}
		if sum == x {
			perfect = append(perfect, x)
		}
	}

	return perfect

}

func divisors(num int) []int {
	sl := []int{}
	for x := 1; x < num; x++ {
		if num%x == 0 {
			sl = append(sl, x)
		}

	}
	return sl
}
