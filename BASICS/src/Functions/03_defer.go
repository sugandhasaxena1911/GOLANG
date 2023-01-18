package main

import "fmt"

func main() {

	openfile()
	defer closefile() // it mkaes sure that file is always closed before the program exits
	otherstuff()
}

func openfile() {

	fmt.Println("Hello , file is opened")
	fmt.Println("Hello , file is edited and contents are written")

}
func closefile() {
	fmt.Println("Hello , file is closed now ")
}

func otherstuff() {
	fmt.Println("continue with other processing ")
}
