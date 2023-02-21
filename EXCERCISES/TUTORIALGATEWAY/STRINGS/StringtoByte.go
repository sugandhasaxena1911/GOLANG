// package in strings by byte
package main

import "fmt"

func main() {
	str := "Hello Sugandha"
	//1
	sl := []byte{}
	for _, v := range str {
		sl = append(sl, byte(v)) //v is int32
	}
	fmt.Println(sl)

	//2
	sl2 := []byte(str)
	fmt.Println(sl2)

	//3
	sl3 := []byte{}
	copy(sl3, str)
	fmt.Println(sl)

	//
	fmt.Printf("%c %v %d", str[0], str[0], str[0])
}
