package view

import (
	"errors"
	"fmt"

	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func Crud() {
	defer func() {
		if r := recover(); r != nil {
			utils.ClearScreen()
			fmt.Println("Terjadi kesalahan yang tidak terduga")
			Crud()
		}
	}()

	var input int
	utils.CrudMenu()
	fmt.Scan(&input)

	switch input {
	case 1:

	case 2:

	case 3:

	case 4:

	case 5:

	case 6:

	case 7:

	case 8:

	case 0:
		utils.ClearScreen()
		Home()

	case 99:
		utils.ClearScreen()

	default:
		err := errors.New("")
		utils.Panik(err)

	}
}
