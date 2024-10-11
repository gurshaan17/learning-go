package main

import "fmt"

func main()  {
	gurshaan := User{"Gurshaan","gurshaan@gmail.com",false,18}
	fmt.Printf("details: %+v\n", gurshaan)
	fmt.Println(gurshaan.Email)
	gurshaan.getstatus()
	gurshaan.changeMail()
	fmt.Println(gurshaan.Email)
}


type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func (u User) getstatus(){
	fmt.Println("User active: ",u.Status)
}

// this only changes value inside the function it isn't pass by reference 
func (u User) changeMail(){
	u.Email = "hello@gmail.com"
	fmt.Println("email of user is: ",u.Email)
}