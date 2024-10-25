package service

import (
	"errors"
	"fmt"

	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
	view "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/View"
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
	view.HomeMenu()
	fmt.Scan(&input)

	switch input {
	case 1:
		LoginMenu()

	case 2:
		DaftarMenu()

	case 3:
		LupaUsername()

	case 4:
		LupaPassword()

	case 99:
		utils.ClearScreen()

	default:
		err := errors.New("")
		utils.Panik(err)
	}

}
