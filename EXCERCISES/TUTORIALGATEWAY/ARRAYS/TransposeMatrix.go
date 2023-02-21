// package in Transose matrix
package main

import "fmt"

func main() {
	rows := 3
	cols := 3
	arr1 := EnterMatrix2(rows, cols)
	//arr2 := TransposeMatrix(arr1)
	fmt.Println(arr1)
	//fmt.Println("Transposed matrix:")
	//fmt.Println(arr2)
	fmt.Println(IsSymmetric(arr1))
}

// Funstion to check symmetrty
func IsSymmetric(arr1 [][]int) bool {
	var b bool
	b = true
	arr2 := TransposeMatrix(arr1)
	fmt.Println(arr1)
	fmt.Println("Transpose of matrix ")
	fmt.Println(arr2)
	r1 := len(arr1)
	c1 := len(arr1[0])

	if r1 != c1 {
		b = false
		return b
	}
	for x := 0; x < r1; x++ {
		for y := 0; y < c1; y++ {
			if arr1[x][y] != arr2[x][y] {
				b = false
				return b
			}
		}
	}

	return b

}
func TransposeMatrix(arr1 [][]int) [][]int {
	// calculate rows and columns of new matrix
	rows := len(arr1[0])
	cols := len(arr1)

	arr2 := [][]int{}
	for r := 0; r < rows; r++ { //3
		arr3 := []int{}
		for c := 0; c < cols; c++ { //2
			arr3 = append(arr3, arr1[c][r])
		}
		arr2 = append(arr2, arr3)
	}

	return arr2
}
func EnterMatrix2(n, m int) [][]int {
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
