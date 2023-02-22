package main

import (
	"fmt"
	"github.com/sugandhasaxena1911/GOLANG/EXCERCISES/TUTORIALGATEWAY/ARRAYS/Operations"
)

func main() {
	fmt.Println("Hello")
	arr1 := Operations.EnterMatrix(3, 3)
	fmt.Println(arr1)
	Operations.Print2DArray(arr1)
	sl := []int{1, 2, 5, 6, 2, 8, 5, 9, 1}
	dupCnt := Operations.CountDuplicates(sl)
	fmt.Println("the number of duplicates were ", dupCnt)
	arr2 := Operations.TransposeMatrix(arr1)
	fmt.Println("After transpose", arr2)

}
