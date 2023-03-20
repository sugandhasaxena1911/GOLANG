package main

import (
	"fmt"
	"github.com/sugandhasaxena1911/GOLANG/EXCERCISES/TUTORIALGATEWAY/NUMBERS/MathOperations"
)

func main() {
	fmt.Println("Hello checking new package ")
	n := 245
	m := MathOperations.NoOfDigits(n)
	fmt.Println("No of digits ", m)
	sl := MathOperations.PerfectNums(300)
	fmt.Println(sl)
}
