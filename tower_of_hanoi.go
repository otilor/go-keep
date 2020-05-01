package main

import (
	"fmt"
)

// Three rods and n number of disks.
// All disks are added to first rod (the mostleft most one)


func TOHUtil(num int, from string, to string, temp string){
	if num < 1{
		return
	}
	TOHUtil(num - 1, from, temp, to)
	fmt.Println("Move disk", num, "from peg", from, "to peg", to)
	TOHUtil(num-1, temp, to, from)
}

func towers_of_hanoi(num int){
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are:")
	TOHUtil(num, "A", "C", "B")
}