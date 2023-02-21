package main

import "fmt"

func main() {
	//size of the square matrix
	//n := 5
	arr1 := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	Printarr1(arr1)
	arr2 := LowerTri(arr1)
	fmt.Println("")
	Printarr1(arr2)

	// because arrays was changing after calling lowertri func()
	arr1 = [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	arr3 := UpperTri(arr1)
	fmt.Println("")
	Printarr1(arr3)
}

func LowerTri(arr1 [][]int) [][]int {
	n := len(arr1)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if y > x {
				arr1[x][y] = 0
			}
		}
	}

	return arr1
}
func UpperTri(arr1 [][]int) [][]int {
	n := len(arr1)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if y < x {
				arr1[x][y] = 0
			}
		}
	}

	return arr1
}

func Printarr1(arr1 [][]int) {
	n := len(arr1)

	for x := 0; x < n; x++ {
		for y := 0; y < len(arr1[x]); y++ {
			// calculate the length of numbers
			dg := 0
			num := arr1[x][y]
			if num == 0 {
				dg = 1
			}
			for num > 0 {
				dg++
				num = num / 10
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
