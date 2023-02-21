package main

import "fmt"

func main() {

	var a byte
	a = 108 // uint8  0-255
	fmt.Println(a, string(a))

	// slice of byte
	b := []byte{101, 103, 69, 79, 55, 54, 154, 97, 99, 108, 65}
	fmt.Println(b, string(b))
	fmt.Printf("%O %o", b[1], b[1])
}
