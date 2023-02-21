package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 7, 56, 2, 89, 6, 7, 56, 1}
	fmt.Println("Number of duplicates ", CountDuplicates(arr))
}

func CountDuplicates(sl []int) int {

	m1 := make(map[int]int, len(sl))
	for _, v := range sl {
		m1[v] = v
	}
	return len(sl) - len(m1)
}
