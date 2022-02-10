package main

import "fmt"

func main() {
	//Variable "a" is manually set to the string "initial", then print "a".
	var a = "initial"
	fmt.Println(a)
	//Variables "b" and "c" are set to integers 1 and 2, respectively, then print "b" and "c".
	var b, c int = 1, 2
	fmt.Println(b, c)
	//Variable "d" is set to the boolean value "true", then print "d".
	var d = true
	fmt.Println(d)
	//Variable "e" is set as an integer with no value. The printed value of "e" is defaulted to 0.
	var e int
	fmt.Println(e)
	//":=" is the shorthand term of doing "var f string = "apple".
	f := "apple"
	fmt.Println(f)
}
