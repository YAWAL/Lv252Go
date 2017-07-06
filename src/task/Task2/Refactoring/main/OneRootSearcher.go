package main

import "fmt"


//TODO refactor this -> add return type!!!
func SearchOneRootInt(n int) (xValue, yValue int) {
	n = 57677
	var x, y int
	//rootCounter := 0
	for x = 1; x < n; x ++ {
		for y = 1; y < n ; y ++  {
			if x*x + y*y == n && y <= x{//
				//rootCounter++
				x = xValue
				y = yValue

				//fmt.Println("Task a -> Pair of natural numbers for n =" , n, "when x^2 + y^2 = n : x =", x, " y =", y)
			}
		}
	}
	//if rootCounter == 0{
	//	fmt.Println("there is no root for n =", n)
	//}
	fmt.Println("egrhrthtrdh",xValue, yValue)
	return xValue, yValue
}
//answer for task 243а page 104
func SearchOneRootBool(n int) bool{
	var x, y int
	rootCounter := 0
	for x = 1; x < n; x ++ {
		for y = 1; y < n ; y ++  {
			if x*x + y*y == n && y <= x{
				rootCounter++
				fmt.Println("Task a -> Pair of natural numbers for n =" , n, "when x^2 + y^2 = n : x =", x, " y =", y)
			}
			if rootCounter > 1 {
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

func main()  {
	fmt.Println(SearchOneRootBool(57677))
	//fmt.Println(SearchOneRootInt(57677))

}

