package service

import (
	"errors"
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func DaftarAkun() (model.Akun, []model.Akun, error) {
	var username, password string

	fmt.Print("Masukkan Username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&password)

	if err := validasiAkun(username, password); err != nil {
		return model.Akun{}, model.DataAkun, err
	}

	model.DataAkun = append(model.DataAkun, model.Akun{Username: username, Password: password})

	if err := utils.SaveAkun("Model/data_akun.json"); err != nil {
		return model.Akun{}, model.DataAkun, err
	}

	return model.Akun{Username: username, Password: password}, model.DataAkun, nil
}

func validasiAkun(username, password string) error {
	for _, akun := range model.DataAkun {
		if akun.Username == username {
			return errors.New("username sudah digunakan")
		}
	}
	if len(username) < 3 {
		return errors.New("username minimal 3 karakter")
	}
	if len(password) < 6 {
		return errors.New("password minimal 6 karakter")
	}

	return nil
}
