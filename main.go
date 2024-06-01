package main

import (
	"fmt"
	"math"
)

func climb_rec(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return climb_rec(n-1) + climb_rec(n-2)
}
func main(){

fmt.Println(climb_rec(5))
}
