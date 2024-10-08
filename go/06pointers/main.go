package main

import "fmt"

func main()  {
	// var ptr *int;
	// fmt.Println(ptr)

	mynum := 34
	var ptr = &mynum
	fmt.Println(ptr) 
}