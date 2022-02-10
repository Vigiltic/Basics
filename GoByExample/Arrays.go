package main

import "fmt"

func main() {
	//a is declared as an array with 5 integers. Since those integers haven't been declared, they are defaulted to 0.
	var a [5]int
	fmt.Println("emp:", a)

	//the fifth integer in array "a" is set to 100. (Note: it counts starting from zero)
	a[4] = 100
	fmt.Println("set:", a)    //Get the entire array printed
	fmt.Println("get:", a[4]) //Get only the fifth integer printed

	//Print the length of *how many* integers are in the array.
	fmt.Println("len:", len(a))

	//Shorthand "b" to be a five-length integer array with the integers set to 1, 2, 3, 4, 5.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//Create two arrays with 3 integers in the first array, and 3 integers in the third array.
	//Using a for loop allows you to access the values of each value of the arrays.
	//The outer array is 2 arrays in length. The inner arrays are 3 integers in length each.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
