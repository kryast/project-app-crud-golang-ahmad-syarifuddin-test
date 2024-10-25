package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func LupaUsername() {
	utils.ClearScreen()
	fmt.Println("---------------------")
	fmt.Println("Daftar Akun saat ini:")

	if err := utils.LoadAkun("Model/data_akun.json"); err != nil {
		fmt.Println("Gagal memuat data akun:", err)
		BackHome()
		return
	}

	for _, a := range model.DataAkun {
		fmt.Printf("Username: %s\n", a.Username)
	}
	fmt.Println("---------------------")
	BackHome()
}
