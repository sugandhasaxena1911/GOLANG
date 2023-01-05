package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d
	)
	//notice how iota + 3 works  : have not hidden this golanD hint
	const (
		e = iota + 3
		f
		g
	)
	const (
		h = iota + 3
		i = iota
		j = iota
	)

	fmt.Printf("the values are %v \t %T", a, a)
	fmt.Printf("the values are %v \t %T", b, b)
	fmt.Printf("the values are %v \t %T", c, c)
	fmt.Printf("the values are %v \t %T", d, d)
	fmt.Printf("the values are %v \t %T", g, g)

}
