package main

import "fmt"

func main() {
	sum := 0
	sl := []int{1, 2, 3, 45, 5}
	for _, v := range sl {

		sum = sum + v
	}
	fmt.Println("The average is ", float64(sum)/float64(len(sl)))
}
