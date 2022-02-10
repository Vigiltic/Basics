package main

import (
	"fmt"
	"time"
)

func main() {
	//i is equal to 2. First, print out "Write 2 as" and see which case it fits. In the case i equals 2, add "two" to the print statement.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	//Take today's day. If it is a Saturday or Sunday, print "It's the weekend". Otherwise, default to printing "It's a weekday".
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	//Take the current time. If the time is before 12 (12:00 PM), then print "It's before noon". Otherwise, default to printing "It's after noon".
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	/*
		Interfaces are a set of types, in this case, bool and int.
		whatAmI is defined below the function with different types.
		"t" takes each type one at a time and compares it with the cases.
		If "t" receives a type that is not a bool or an int, it'll print the default case, which will print "don't know type 'type' (in this case, a string)"
	*/
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
