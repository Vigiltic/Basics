package main

import "fmt"

func main() {
	//If 7 divided by 2 has a remainder of 0, print "7 is even", else print "7 is odd"
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	//if statement without an else statement. If there is no else statement, and if the "if" statement is false, then nothing is printed.
	if 9%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	//A statement can precede conditionals; any variables declared in this statement are available in all branches.
	/*
		If the given value is less than 0, print the first conditional. If it is less than 10, print the second conditional.
		If the other two statements were unfulfilled, print the very last conditional.
	*/
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
