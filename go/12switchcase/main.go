package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Seed(time.Now().Unix())
	//this line was required in older version of go now it is not required

	dice := rand.Intn(6) + 1

	fmt.Println("value of dice is:",dice)

	switch dice {
	case 1:
		fmt.Println("value of dice is 1")
	case 2:
		fmt.Println("value of dice is 2")
	case 3:
		fmt.Println("value of dice is 3")
	case 4:
		fmt.Println("value of dice is 4")
		fallthrough
	case 5:
		fmt.Println("value of dice is 5")
		fallthrough
	case 6:
		fmt.Println("value of dice is 6")
	}


}