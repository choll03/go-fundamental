package main

import "fmt"

func main() {

	// 1. deklarasi map
	var chiken = map[string]int{
		"januari": 10,
		"maret":   5,
		"april":   9,
	}

	looping(chiken)

	// 3. Menghapus Item Map
	delete(chiken, "januari")

	looping(chiken)

	// 4. Deteksi keberadaan item dengan key tertentu
	var value, isExist = chiken["tes"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item not exists")
	}

	// 5. Kombinasi slice & map
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken pink", "gender": "female"},
	}

	for key, chicken := range chickens {
		fmt.Println(key, chicken["name"])
		fmt.Println(key, chicken["gender"])
	}
}

// 2. Iterasi Item Map Menggunakan for - range
func looping(chicken map[string]int) {
	fmt.Println("----------------------------------")

	for key, val := range chicken {
		fmt.Println(key, val)
	}

}
