package main

import "fmt"

func main()  {
	var fruits [4]string;
	fruits[0] = "apple"
	fruits[2]=  "banana"
	fruits[3] = "kiwi"

	fmt.Println(fruits)
	// [apple  banana kiwi]
	//it will print this the empty space indicates there is an entry missing inside the array

	var veg = [3]string {"potato","tomato","onion"}

	fmt.Println(veg)
	fmt.Println(len(veg))
}