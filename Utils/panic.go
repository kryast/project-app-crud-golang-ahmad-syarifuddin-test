package utils

import "fmt"

func Panik(err error) {
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
		panic(err)
	}
}
