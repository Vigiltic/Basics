package main

import "fmt"

func main() {
	//Make slice "s" with 3 empty string values.
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"                //Set the first value of slice "s" to "a"
	s[1] = "b"                //Set the second value of slice "s" to "b"
	s[2] = "c"                //Set the third value of slice "s" to "c"
	fmt.Println("set:", s)    //Print all the values of s
	fmt.Println("get:", s[2]) //Print the third value of s

	fmt.Println("len:", len(s)) //Print the length of s

	//Using append allows you to add another value to a slice. Note this will not work with an array.
	s = append(s, "d")      //Adding one value here.
	s = append(s, "e", "f") //Adding two values here.
	fmt.Println("apd:", s)

	//Making a new slice which has the length of slice s.
	c := make([]string, len(s))
	copy(c, s) //Copying the values of s into the values of c [Recipient, Original]
	fmt.Println("cpy:", c)

	//A slice operator which gets the values of the third value up to the fifth value.
	//Note that the operators' inclusion involves an [including, excluding] formula.
	l := s[2:5]
	fmt.Println("sl1:", l)

	//A slice operator which gets all values of s, excluding the sixth (and any beyond) digit.
	l = s[:5]
	fmt.Println("sl2:", l)

	//Starts the slice operator after and including the third value.
	l = s[2:]
	fmt.Println("sl3:", l)

	//Declare and initialize a slice (using an empty []) with the values "g", "h", and "i".
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//Create two integer slices with a maximum length of 3.
	//The slices cannot exceed a length of 3, but can have less than 3 values.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
