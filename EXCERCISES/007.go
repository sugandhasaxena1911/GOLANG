package main

import "fmt"

/*Write a program which takes 2 digits, X,Y as input and generates a 2-dimensional array.
The element value in the i-th row and j-th column of the array should be i*j.
Note: i=0,1.., X-1; j=0,1,¡­Y-1.

Example

Suppose the following inputs are given to the program: 3,5

Then, the output of the program should be: [[0, 0, 0, 0, 0], [0, 1, 2, 3, 4], [0, 2, 4, 6, 8]]
*/

func main() {
	num1 := 3
	num2 := 5

	arr := [][]int{}
	for x := 0; x < num1; x++ {
		temp := []int{}
		for y := 0; y < num2; y++ {
			temp = append(temp, x*y)
		}
		arr = append(arr, temp)
	}

	fmt.Println(arr)

	//pratice unfurling on 2d array
	arr2 := [][]int{}
	arr2 = append(arr2, arr...)
	fmt.Println(arr2)

}
