package view

import "fmt"

func HomeMenu() {
	fmt.Println("================================================")
	fmt.Println("Selamat Datang Di SMA YPI Tunas Bangsa Palembang")
	fmt.Println("================================================")
	fmt.Println("1. Login")
	fmt.Println("2. Daftar")
	fmt.Println("3. Lupa Username")
	fmt.Println("4. Lupa Password")
	fmt.Println("99. Exit")
	fmt.Print("Masukkan nomor menu: ")
}
