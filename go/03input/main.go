package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	fmt.Println("welcome to input program")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter anything: ")

	input , _ := reader.ReadString('\n')
	fmt.Print("You entered: ",input)

}