package main

import "fmt"

func main() {
	i := 1
	//For the case of i being unequal to 3, print the value of i, add 1 to it, then print again until i equals 3.
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	/*
		j is set to 7. For the case of j being unequal to 9, print the value of j, add 1 to it (with the "j++" command),
		then print again until j equals 9.
	*/
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//An infinite loop. Using the 'break' feature will "break out" of the loop and continue on with the rest of the code.
	for {
		fmt.Println("loop")
		break
	}
	//This loop will print the value of n if there is any remainder after n is divided by 2. If there is no remainder, nothing is printed.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
