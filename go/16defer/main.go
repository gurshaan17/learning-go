package main

import "fmt"

func main(){
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("hello world")
	my()
}

func my() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}