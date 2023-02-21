package main

import "fmt"

func main() {

	//size of matrix = n
	n := 3
	arr1 := [][]int{{1, 0, 0}, {0, 1, 0}, {1, 0, 1}}
	b := CheckIdentity(arr1)
	fmt.Println("The matrix of size ", n, "is a identity : ", b)
}
func CheckIdentity(arr1 [][]int) bool {
	//var identity bool
	identity := true
	//size of square matrix
	n := len(arr1)
	for x := 0; x < n; x++ {

		for y := 0; y < n; y++ {
			if y == x {
				if arr1[x][y] != 1 {
					identity = false
					break
				}
			} else {
				if arr1[x][y] != 0 {
					identity = false
					break
				}
			}

		}

	}
	return identity

}
