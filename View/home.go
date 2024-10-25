package view

import (
	"errors"
	"fmt"

	service "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Service"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func Home() {
	defer func() {
		if r := recover(); r != nil {
			utils.ClearScreen()
			fmt.Println("Terjadi kesalahan yang tidak terduga")
			Home()
		}
	}()

	var input int
	utils.HomeMenu()
	fmt.Scan(&input)

	switch input {

	case 2:
		utils.ClearScreen()
		fmt.Println("Silahkan Daftar Akun")
		akun, dataAkun, err := service.DaftarAkun()
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

		Home()

	case 99:
		utils.ClearScreen()

	default:
		err := errors.New("")
		utils.Panik(err)
	}

}
