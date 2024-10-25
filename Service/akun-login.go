package service

import (
	"errors"
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func Login() error {
	var username, password string

	if err := utils.LoadAkun("Model/data_akun.json"); err != nil {
		return errors.New("gagal memuat data akun: " + err.Error())
	}

	fmt.Print("Masukkan Username: ")
	fmt.Scan(&username)

	for _, akun := range model.DataAkun {
		if akun.Username == username {
			fmt.Print("Masukkan Password: ")
			fmt.Scan(&password)
			if akun.Password == password {
				fmt.Println("Berhasil Login!")
				return nil
			}
			return errors.New("password salah")
		}
	}

	return errors.New("username tidak ditemukan, silakan daftar terlebih dahulu")
}
