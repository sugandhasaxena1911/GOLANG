package main

import "fmt"

func main() {

	sl := []int{1, 3, 7, 0, 98, 4, 678}
	min, max := MinMax(sl)
	fmt.Println("The mix .max are ", min, max)
}

func MinMax(sl []int) (int, int) {
	min := sl[0]
	max := sl[0]
	for _, v := range sl {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}

	}
	return min, max
}
