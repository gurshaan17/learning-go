package main

import "fmt"


var MyUser  = "New User"


func main() {
	var username string = "tom_18"
	fmt.Println(username)
	fmt.Printf("variable type: %T \n",username)

	var smallInt uint8 = 234
	fmt.Println(smallInt)

	var isTrue bool = true
	fmt.Println(isTrue)

	var smallFloat float32 = 24234.23423432
	fmt.Print(smallFloat,"\n")

	// we can also directly declare variables
	var usr = "gurshaan"
	fmt.Print(usr,"\n")

	//walrus operator
	usrnm:="gurshaan17"
	fmt.Print(usrnm)

	fmt.Println("\n",MyUser)
}

