package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var str string
	fmt.Println("Enter comma separated numbers")
	_, err := fmt.Scanln(&str)
	if err != nil {
		log.Fatal("Cannot proceed", err)
	}

	st := strings.Split(str, ",")
	fmt.Println(st, len(st))
	intsl := to_int_slice(st)
	fmt.Printf("%T %v", intsl, intsl)

}

func to_int_slice(str []string) []int {
	// start parsing the string
	num := make([]int, len(str))
	for i, v := range str {
		n, _ := strconv.Atoi(v)
		num[i] = n
	}
	return num
}
