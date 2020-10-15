package main

import "fmt"

func main() {

	var firsName string = "jonh"

	var lastName string
	lastName = "wick"

	fmt.Printf("helo %s %s!\n", firsName, lastName)

	name := new(string)

	fmt.Println(name)
	fmt.Println(*name)
}
