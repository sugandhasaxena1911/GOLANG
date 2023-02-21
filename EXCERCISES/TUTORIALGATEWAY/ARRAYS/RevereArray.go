package main

import "fmt"

func main() {
	arr := []int{1, 7, 3, 7, 8}

	sl := ReverseArray(arr)
	fmt.Println(sl)

}
func ReverseArray(sl []int) []int {
	m := len(sl)
	for i, _ := range sl {
		if i >= m/2 {
			break
		}
		sl[i], sl[m-i-1] = sl[m-i-1], sl[i]

	}

	return sl
}
