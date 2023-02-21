package main

import (
	"fmt"
	"math"
)

func main() {

	//these are all constants

	fmt.Println("the MAX value of INT is ", math.MaxInt) //64 bit bcause processor is 64 bit
	fmt.Println("the MIN value of INT is ", math.MinInt)

	fmt.Println("the MAX value of INT8 is ", math.MaxInt8) // 1<<7  -1
	fmt.Println("the MIN value of INT8 is ", math.MinInt8) //-1<<7

	fmt.Println("the MAX value of INT16 is ", math.MaxInt16) // 1<<15 -1
	fmt.Println("the MIN value of INT16 is ", math.MinInt16) //1<<15

	fmt.Println("the MAX value of INT32 is ", math.MaxInt32)
	fmt.Println("the MIN value of INT32 is ", math.MinInt32)

	fmt.Println("the MAX value of INT64 is ", math.MaxInt64)
	fmt.Println("the MIN value of INT64 is ", math.MinInt64)

	//UNSIGNED, no max because its alays 0
	fmt.Println("the MAX value of UINT8 is ", math.MaxUint8)   //1<<8 -1
	fmt.Println("the MIN value of UINT16 is ", math.MaxUint16) //1<<16-1
	fmt.Println("the MAX value of UINT32 is ", math.MaxUint32) //1<<32-1
	//fmt.Println("the MIN value of INT64 is ", math.MaxUint64)  //1<<64 -1   //overflows

}
