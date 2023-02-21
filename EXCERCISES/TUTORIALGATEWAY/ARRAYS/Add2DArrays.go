package main

import "fmt"

func main() {
	//5*3
	arr1 := [][]int{{1, 3, 4}, {5, 7, 8}, {16, 5, 2}, {1, 2, 3}, {6, 4, 3}}
	arr2 := [][]int{{1, 3, 4}, {5, 7, 8}, {16, 5, 2}, {1, 2, 3}, {6, 4, 3}}
	arr3 := Add2DArray(arr1, arr2)
	Print2DArray(arr3)

}
func Add2DArray(arr1, arr2 [][]int) [][]int {
	arr3 := [][]int{}
	for i := 0; i < len(arr1); i++ {
		temp := []int{}
		for j := 0; j < len(arr1[i]); j++ {
			temp = append(temp, arr1[i][j]+arr2[i][j])
		}
		arr3 = append(arr3, temp)
	}
	return arr3
}

func Print2DArray(arr1 [][]int) {
	for i, _ := range arr1 {
		for _, vv := range arr1[i] {
			fmt.Print(vv, " ")

		}
		fmt.Println(" ")

	}
}
