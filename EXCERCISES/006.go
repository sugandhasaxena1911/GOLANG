package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Q = Square root of [(2 * C * D)/H]
// c,H const  50,30 and D = comman separted numbers
// INPUT : 100,150,180
//
// The output of the program should be: 18,22,24
const (
	c = 50
	h = 30
)

func main() {
	var str string
	fmt.Println("Enter comma separted int numbers ")
	_, err := fmt.Scanln(&str)
	if err != nil {
		log.Fatalln("cannot proceed ", err)
	}
	st := strings.Split(str, ",")
	fmt.Println(st)
	fmt.Println("The result is ", sqrt(st))

}

func sqrt(st []string) []float64 {
	slnum := make([]float64, len(st))
	//var r float64

	for i, v := range st {
		n, _ := strconv.Atoi(v)

		r := (float64(2*c*n) / h) /// very crucial step where u need to specify float64 s, otherwise output comes as int
		fmt.Println(r)
		slnum[i] = math.Sqrt((r))

		//if you want slice of ints, make slice of ints for slnum and then use Round
		//math.Round(math.Sqrt((r)))
	}

	return slnum
}
