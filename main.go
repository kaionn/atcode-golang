package main

import (
	"fmt"
)

func correct_numbers(a, b, c int) bool {
	return a < c && c < b
}

func main() {
	var n [3]int
	fmt.Scan(&n[0], &n[1], &n[2])
	if correct_numbers(n[0], n[1], n[2]) {
		fmt.Println("Yes")
	} else {
		fmt.Println("False")
	}
}