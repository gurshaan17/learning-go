package main

import "fmt"

func main()  {

	gurshaan := User{"Gurshaan","gurshaan@gmail.com",false,18}
	fmt.Printf("details: %+v\n", gurshaan)
	fmt.Println(gurshaan.Email)
}


type User struct {
	Name string
	Email string
	Status bool
	Age int
}