package main

import "fmt"

/*this function search one root of equition n = x^2 + y^2
function accept int value of n and return bool value
*/
func SearchOneRoot(n int) bool{
	var x, y int//counters for loops
	rootCounter := 0
	for x = 1; x < n; x ++ {
		for y = 1; y < n ; y ++  {
			if x*x + y*y == n {
				rootCounter++
				fmt.Println("Task a -> Pair of natural numbers for n =" , n, "when x^2 + y^2 = n : x =", x, " y =", y)
			}
			if rootCounter >= 1 {
				break
			}
		}
	}
	if rootCounter == 0{
		fmt.Println("there is no root for n =", n)
		return false
	}
	return true
}