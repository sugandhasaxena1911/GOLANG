package main

/*
Create a seperate file (module) which has at least two methods:
- ReadString: to read a string from console input
- PrintString: to return the string in upper case.
Also create a `main.go` file that acts as calling class.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input string = ""

func main() {
	ReadString()
	fmt.Println(PrintString())
}

// ReadString reads user input from command line, terminated via 'enter'
func ReadString() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
}

// PrintString prints the string entered via GetString
func PrintString() string {
	return strings.ToUpper(input)
}
