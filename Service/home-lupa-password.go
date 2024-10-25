package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func LupaPassword() {
	utils.ClearScreen()

	if err := utils.LoadAkun("Model/data_akun.json"); err != nil {
		fmt.Println("Gagal memuat data akun:", err)
		BackHome()
		return
	}

	if len(model.DataAkun) == 0 {
		fmt.Println("Data Masih kosong")
		BackHome()
		return
	}

	var input string
	fmt.Print("Masukkan Username: ")
	fmt.Scan(&input)

	found := false

	for _, akun := range model.DataAkun {
		if akun.Username == input {
			fmt.Printf("Username: %s\nPassword: %s\n", akun.Username, akun.Password)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Error: Username tidak ditemukan.")
	}
	BackHome()
}
