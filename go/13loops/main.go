package main

import "fmt"

//7362
func main()  {
	days := []string {"monday","tuesday","wednesday"}

	for d:=0;d<len(days);d++{
		fmt.Println(days[d])
	}

	for i := range(days){
		fmt.Println(days[i])
	}

	for i,day := range(days) {
		fmt.Println(day)
		fmt.Println(days[i])
	}

	val := 1
	for val < 10 {

		if val == 6 {
			goto greet
		}

		if val == 3 {
			val++
			continue
		}

		if val == 7 {
			break
		}

		fmt.Println(val)
		val++
	}

	greet:
		fmt.Println("hello")
}