package service

import (
	"errors"
	"fmt"

	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
	view "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/View"
)

func Crud() {
	defer func() {
		if r := recover(); r != nil {
			utils.ClearScreen()
			fmt.Println("Terjadi kesalahan yang tidak terduga")
			Crud()
		}
	}()

	if err := utils.LoadGuru("Model/data_guru.json"); err != nil {
		fmt.Println("Error loading guru data:", err)
	}
	if err := utils.LoadSiswa("Model/data_siswa.json"); err != nil {
		fmt.Println("Error loading siswa data:", err)
	}

	var input int
	view.CrudMenu()
	fmt.Scan(&input)

	switch input {

	case 1:
		CreateGuru()

	case 2:
		CreateSiswa()

	case 3:
		UpdateGuru()

	case 4:
		UpdateSiswa()

	case 5:
		DeleteGuru()

	case 6:
		DeleteSiswa()

	case 7:
		ReadGuru()

	case 8:
		ReadSiswa()

	case 0:
		utils.ClearScreen()
		Home()

	case 99:
		utils.ClearScreen()
		fmt.Println("Terima kasih! Sampai jumpa!")
		return

	default:
		err := errors.New("pilihan tidak valid")
		utils.Panik(err)
	}

	Crud()
}
