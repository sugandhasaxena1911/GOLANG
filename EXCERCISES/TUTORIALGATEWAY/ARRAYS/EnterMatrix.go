package main

import "fmt"

func main() {
	arr1 := Entermatrix(2, 3)
	fmt.Println(arr1)
}

func Entermatrix(n, m int) [][]int {
	var num int

	arr1 := [][]int{}
	for x := 0; x < n; x++ {
		arr2 := []int{}
		for y := 0; y < m; y++ {
			fmt.Print("Enter number")
			fmt.Scan(&num)
			arr2 = append(arr2, num)
		}
		arr1 = append(arr1, arr2)
	}

	return arr1

}
