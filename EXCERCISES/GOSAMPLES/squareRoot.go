package main

import (
	"fmt"
	"math"
)

func main() {

	var a int
	a = 25
	//b := 25
	c := 25.0

	//gives error : cannot use a (variable of type int) as type float64 i
	//fmt.Println("the square root of 25 is ", math.Sqrt(a))
	fmt.Println("the square root of 25 is ", math.Sqrt(float64(a)))
	fmt.Println("the square root of 25 is ", math.Sqrt(25))
	//fmt.Println("the square root of 25 is ", math.Sqrt(b))  //gives same erorr as above
	fmt.Println("the square root of 25 is ", math.Sqrt(c))

}
