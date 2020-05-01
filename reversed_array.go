package main

import (
	"fmt"
	
)

func reverse(a[]int){
	for j:= 0; j < len(a); j ++{
		fmt.Println(a[len(a)-1-j])
	}
	
}

