package main

import (
	"fmt"
)

func reverse_string(s string) {
	var a string = "abcd"
	var reversed []int
	

	for i := range a {
		reversed = append(reversed, i)
		fmt.Printf("%c", reversed)
	}
}

