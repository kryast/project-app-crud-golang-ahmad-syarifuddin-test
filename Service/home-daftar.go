package service

import (
	"fmt"

	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func DaftarMenu() {
	utils.ClearScreen()
	fmt.Println("Silahkan Daftar Akun")
	akun, dataAkun, err := DaftarAkun()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("---------------------")
		fmt.Println("Akun berhasil dibuat!")
		fmt.Printf("Username: %s\n", akun.Username)
		fmt.Println("---------------------")
		fmt.Println("Daftar Akun saat ini:")
		for _, a := range dataAkun {
			fmt.Printf("Username: %s\n", a.Username)
		}
		fmt.Println("---------------------")
	}

	BackHome()
}
