package service

import (
	"errors"
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
)

func Login() error {
	var username string
	var password string

	for _, akun := range model.DataAkun {
		fmt.Print("Masukkan Username: ")
		fmt.Scan(&username)
		if akun.Username == username {
			fmt.Print("Masukkan Password: ")
			fmt.Scan(&password)
			if akun.Password == password {
				fmt.Println("Berhasil Login!")
				return nil
			}
			return errors.New("password salah")
		} else {

			return errors.New("username salah")
		}
	}
	return errors.New("tidak ada username yang terdaftar. silahkan daftar terlebih dahulu")
}
