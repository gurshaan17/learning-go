package main

import (
	"fmt"
	"sort"
)

func main()  {
	var fruits = []string {}
	fruits = append(fruits, "Apple", "Banana", "Mango", "Kiwi")

	fmt.Println(fruits)

	// fruits = append(fruits[1:3])

	fmt.Println(fruits)

	high := make([]int,4)

	high[0]=23
	high[1]=33
	high[2]=12
	high = append(high, 4, 45)

	fmt.Println(high)

	fmt.Println(sort.IntsAreSorted(high))

	sort.Ints(high)

	fmt.Println(high)

	fmt.Println(sort.IntsAreSorted(high))



	var courses = []string{"java","js","php","cpp"}
	fmt.Println(courses)

	//this index will get removed from the slice
	var index int = 2;
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)


}