package main

import "fmt"

func main() {
	fmt.Println("functions in golang")
	greet()
	result := adder(3,5)
	fmt.Println(result)
	res := pro(2,6,7)
	fmt.Println(res)
	o1,o2 := pro2(7,8,0,9,56)
	fmt.Println(o1)
	fmt.Println(o2)
}

func greet(){
	fmt.Println("hello")
}

func adder(val1 int,val2 int)int{
	return val1+val2
}



func pro (value ...int) int{
	total := 0
	for _,val := range value{
		total += val
	}
	return total
}
func pro2 (value ...int) (int,string){
	total := 0
	for _,val := range value{
		total += val
	}
	return total,"hi pro "
}