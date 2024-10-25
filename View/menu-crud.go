package view

import "fmt"

func CrudMenu() {
	fmt.Println("================================================")
	fmt.Println("Selamat Datang Di SMA YPI Tunas Bangsa Palembang")
	fmt.Println("================================================")
	fmt.Println("1. Input Data Guru")
	fmt.Println("2. Input Data Siswa")
	fmt.Println("3. Update Data Guru")
	fmt.Println("4. Update Data Siswa")
	fmt.Println("5. Delete Data Guru")
	fmt.Println("6. Delete Data Siswa")
	fmt.Println("7. Tampilkan Semua Data Guru")
	fmt.Println("8. Tampilkan Semua Data Siswa")
	fmt.Println("0. Log Out")
	fmt.Print("Masukkan nomor menu: ")
}
