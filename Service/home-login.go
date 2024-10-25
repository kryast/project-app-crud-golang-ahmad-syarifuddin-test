package service

import (
	"fmt"

	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func LoginMenu() {

	utils.ClearScreen()
	utils.ClearScreen()
	fmt.Println("Silahkan Login")
	err := Login()
	if err != nil {
		fmt.Println("Error :", err)
		BackHome()
	} else {
		utils.ClearScreen()

		Crud()
	}
}
