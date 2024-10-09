package main

import "fmt"

func main()  {
	languages := make( map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby" 
	languages["py"] = "python"

	fmt.Println(languages)
	fmt.Println(languages["js"])

	delete(languages,"rb")
	fmt.Println(languages)

	for key, value:= range languages {
		fmt.Printf("key: %v value: %v \n",key,value)
	}
}