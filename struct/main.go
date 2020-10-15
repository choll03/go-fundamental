package main

import "fmt"

type Mahasiswa struct {
	nim   string
	nama  string
	kelas string
}

func main() {

	var mahasiswas = []Mahasiswa{
		{nim: "2016141807", nama: "ismail", kelas: "09TPLE005"},
		{nim: "2016141389", nama: "Nuy Kasdut", kelas: "09TPLE006"},
	}

	for _, mahasiswa := range mahasiswas {
		fmt.Println(mahasiswa)
	}
}
