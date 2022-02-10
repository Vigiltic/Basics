package main

import (
	"fmt"
	"math"
)

//Constant "s" is set as the string "constant". This cannot be changed in the package level.
const s string = "constant"

func main() {
	//Print the value of "s".
	fmt.Println(s)
	//New constant "n" is set to 500,000,000
	const n = 500000000
	//New constant "d" is set to a number with E notation
	const d = 3e20 / n
	fmt.Println(d)
	//Constant "d" is converted into int64 for readability
	fmt.Println(int64(d))
	//Constant "n" is converted into a float and taken the Sin from.
	fmt.Println(math.Sin(n))
}
