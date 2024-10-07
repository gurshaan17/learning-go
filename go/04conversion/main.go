package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Print("Welcome to our pizza app, enter rating: ")

	reader := bufio.NewReader(os.Stdin)

	input , _ := reader.ReadString('\n')

	fmt.Print("You entered: ",input)

	num,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating: ",num+1)
	}

}