package main

import "fmt"

func main()  {
	var num  = 2
	if num%2==0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 3; num<10 {
		fmt.Println("number less than 10")
	} else {
		fmt.Println("number greater than 10")
	}
}