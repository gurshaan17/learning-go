package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome to time module of go lang")

	present := time.Now()

	fmt.Println(present)

	frmtDate := present.Format("01-02-2006 Monday")

	fmt.Println(frmtDate)

	frmtTime := present.Format("15:04:05")

	fmt.Println(frmtTime)

	createTime := time.Date(2024,time.November,10,06,25,28,0,time.Local)

	fmt.Println(createTime)
}

