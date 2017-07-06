package main

import "fmt"

//answer for task 243б page 104
func SearchAllRoots(n int) bool {
	var x, y int//counters for loops
	rootCounter := 0
	for x = 1; x < n; x ++ {
		for y = 1; y < n ; y ++  {
			if x*x + y*y == n && y <= x {
				rootCounter++
				fmt.Println("Task б -> Pairs of natural numbers for n =" , n, "when x^2 + y^2 = n and x >= y: x =", x, " y =", y)
			}
		}
	}
	if rootCounter == 0{
		fmt.Println("there is no root for n =", n)
		return false
	}
	return true
}