package main

import (
	"fmt"
	"strings"
)

type FilterCallback func(string) bool

func filters(data []string, callback FilterCallback) []string {

	var result []string

	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}

func main() {
	var data = []string{"is", "mail", "ndut"}

	var dataFilter = filters(data, func(each string) bool {
		return strings.Contains(each, "u")
	})

	fmt.Println("Data asli", data)
	fmt.Println("Data filter", dataFilter)
}
