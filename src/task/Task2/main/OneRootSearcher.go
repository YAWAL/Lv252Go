package main

import "fmt"

//answer for task 243а page 104
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

//func main()  {
//
//	n := 18
//	SearchOneRoot(n)
//
//	fmt.Println(SearchOneRoot(n))
//	//fmt.Println(SearchOneRootInt(57677))
//
//}
