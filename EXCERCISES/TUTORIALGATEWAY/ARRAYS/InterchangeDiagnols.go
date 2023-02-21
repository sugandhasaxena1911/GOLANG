package main

import "fmt"

func main() {
	//size of the square matrix
	//n := 5
	arr1 := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	Printarr(arr1)
	arr2 := Interchange(arr1)
	fmt.Println("")
	Printarr(arr2)
}

func Interchange(arr1 [][]int) [][]int {
	//diag1 : 0,0  1,1   2,2 ..... diag2 : 0,n-1   1,n-1 ....
	n := len(arr1)

	for x := 0; x < n; x++ {

		arr1[x][x], arr1[x][n-x-1] = arr1[x][n-x-1], arr1[x][x]
	}

	return arr1

}

func Printarr(arr1 [][]int) {
	n := len(arr1)

	for x := 0; x < n; x++ {
		for y := 0; y < len(arr1[x]); y++ {
			// calculate the length of numbers
			dg := 0
			num := arr1[x][y]
			for num > 0 {
				num = num / 10
				dg++
			}
			space := 4 - dg
			str := ""
			switch space {
			case 1:
				str = " "
			case 2:
				str = "  "
			case 3:
				str = "   "

			}
			fmt.Print(arr1[x][y], str)
		}
		fmt.Println("")
	}
}
