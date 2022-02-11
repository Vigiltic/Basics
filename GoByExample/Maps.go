package main

import "fmt"

func main() {
	//Maps are Go's built-in associative data type. Line 7 makes an empty map m with the format "string":"integer" for later use.
	m := make(map[string]int)

	//Lines 10 and 11 set the key/value pairs inside map m.
	m["k1"] = 7
	m["k2"] = 13

	//Printed statement will be "map: map[k1:7 k2:13]"
	fmt.Println("map:", m)

	//Line 17 gets the value of k1 and copies it to v1.
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	//This will print out how many value pairs are inside map m.
	fmt.Println("len:", len(m))

	//Pretty self-explanatory. The delete command will remove k2 from the map.
	delete(m, "k2")
	fmt.Println("map:", m)

	//Line 28 checks to see if k2 is still present in map m. A blank identifier "_" is used since we don't need the value itself.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//Declaring and initializing a map in the same line.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
