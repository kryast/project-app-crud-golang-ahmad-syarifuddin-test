package view

import (
	"fmt"

	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func BackCrud() {
	var input int
	fmt.Println("=============================")
	fmt.Println("Silahkan Kembali ke Menu Database")
	fmt.Println("=============================")
	fmt.Println("0. Kembali")
	fmt.Print("Masukkan nomor menu: ")
	fmt.Scan(&input)
	utils.ClearScreen()
	if input == 0 {
		Crud()
	} else {
		fmt.Println("Error !!")
		Crud()
	}
}
