package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Find numbers devisable by 7 but not by 5, 2000 and 3200 and print in single string with comman separetd values
func main() {
	final_str := divnum(2000, 3200)
	fmt.Println("The final output ", final_str)
}

func divnum(lower_l int, upper_l int) string {
	var str []string
	for x := lower_l; x <= upper_l; x++ {
		if x%7 == 0 && x%5 != 0 {
			//fmt.Println(strconv.Itoa(x))
			str = append(str, strconv.Itoa(x))
		}
	}
	return (strings.Join(str, ","))
}
