package view

import (
	"errors"
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
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
	case 1:
		utils.ClearScreen()
		utils.ClearScreen()
		fmt.Println("Silahkan Login")
		err := service.Login()
		if err != nil {
			fmt.Println("Error :", err)
			BackHome()
		} else {
			utils.ClearScreen()

			Crud()
		}

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

		BackHome()

	case 3:
		utils.ClearScreen()
		fmt.Println("---------------------")
		fmt.Println("Daftar Akun saat ini:")

		if err := utils.LoadAkun("data_akun.json"); err != nil {
			fmt.Println("Gagal memuat data akun:", err)
			BackHome()
			return
		}

		for _, a := range model.DataAkun {
			fmt.Printf("Username: %s\n", a.Username)
		}
		fmt.Println("---------------------")
		BackHome()

	case 4:
		utils.ClearScreen()

		if err := utils.LoadAkun("data_akun.json"); err != nil {
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

	case 99:
		utils.ClearScreen()

	default:
		err := errors.New("")
		utils.Panik(err)
	}

}
