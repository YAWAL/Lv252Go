package main

import (
	"fmt"
)

func WriteTask()  {
	fmt.Println("Дано натуральне число n. \n" +
		"Чи можна представити його у вигляді суми двох квадратів натуральних чисел? \n" +
		"Якщо можна, то : \n" +
		"a) вказати пару х, у таких натуральних чисел, що n = х^2 + y^2 \n" +
		"б) вказати всі пари х, у таких натуральних чисел, що n = х^2 + y^2, x >= y \n")
}

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

//this function search all roots of equition n = x^2 + y^2 , when x >= y
//function accept int value of n and return bool value
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

func main()  {

	WriteTask()

	for   {
		var n int
		fmt.Println("Please, enter n : ")
		fmt.Scanln(&n)
		SearchOneRoot(n)
		SearchAllRoots(n)
	}
}
