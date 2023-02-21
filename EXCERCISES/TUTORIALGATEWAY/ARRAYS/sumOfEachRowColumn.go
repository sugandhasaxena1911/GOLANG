// package sum of row columns
package main

import "fmt"

func main() {
	arr1 := Entermatrix1(2, 3)
	fmt.Println(arr1)
	SumOfEachColumnRow(arr1)
}

func Entermatrix1(n, m int) [][]int {
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

func SumOfEachColumnRow(arr1 [][]int) {
	for i, _ := range arr1 {
		sumRow := 0
		for _, vv := range arr1[i] {
			sumRow = sumRow + vv
		}
		fmt.Println("Sum of Row", i, "is : ", sumRow)
	}

	for y := 0; y < len(arr1[0]); y++ {
		sumColumn := 0
		for x := 0; x < len(arr1); x++ {
			sumColumn = sumColumn + arr1[x][y]
		}
		fmt.Println("Sum of column", y, "is : ", sumColumn)

	}
}
