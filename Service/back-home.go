package service

import (
	"errors"
	"fmt"

	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func BackHome() {
	var input int
	fmt.Println("=============================")
	fmt.Println("Silahkan Kembali ke Menu Home")
	fmt.Println("=============================")
	fmt.Println("0. Kembali")
	fmt.Print("Masukkan nomor menu: ")
	fmt.Scan(&input)
	utils.ClearScreen()
	if input == 0 {
		Home()
	} else {
		err := errors.New("")
		utils.Panik(err)
	}
}
