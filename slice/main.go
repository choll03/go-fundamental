package main

import "fmt"

func main() {

	var fruits = []string{"apple", "grape", "banana", "melon"}

	var newFruit = fruits[0:2]

	fmt.Println(newFruit)
	fmt.Println(fruits)
}
