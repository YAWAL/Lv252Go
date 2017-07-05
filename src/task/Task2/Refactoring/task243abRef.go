package main

import "fmt"

func WriteTask()  {
	fmt.Println("Дано натуральне число n. \n" +
		"Чи можна представити його у вигляді суми двох квадратів натуральних чисел? \n" +
		"Якщо можна, то : \n" +
		"a) вказати пару х, у таких натуральних чисел, що n = х^2 + y^2 \n" +
		"б) вказати всі пари х, у таких натуральних чисел, що n = х^2 + y^2, x >= y \n")
}

//answer for task 243а page 104
//TODO refactor this -> add return type!!!
func SearchOneRoot(n int) (xValue, yValue int) {
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



func main()  {

	//WriteTask()
	for   {
		//var n int
		//fmt.Println("Please, enter n : ")
		//fmt.Scanln(&n)
		SearchOneRoot(57677)
	}
}
