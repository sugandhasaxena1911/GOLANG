package main

import (
	"fmt"
	"math"
)

func main() {
	//its a constant defined the package
	fmt.Println("The value of pi is ", math.Pi)        // 3.141592653589793
	fmt.Printf("The value of pi is  %.5f ", math.Pi)   // width default , precison : 5
	fmt.Printf("The value of pi is  %9.55f ", math.Pi) // width 9 , precision 55

}
