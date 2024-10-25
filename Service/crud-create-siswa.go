package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func CreateSiswa() {
	var siswa model.Siswa
	siswa.ID = utils.TambahID(model.DataSiswa)

	fmt.Print("Masukkan nama siswa : ")
	if _, err := fmt.Scan(&siswa.Nama); err != nil || siswa.Nama == "99" {
		utils.ClearScreen()
		Home()
		return
	}

	fmt.Print("Masukkan NIS siswa : ")
	if _, err := fmt.Scan(&siswa.NIS); err != nil || siswa.NIS == 99 {
		utils.ClearScreen()
		Home()
		return
	}

	fmt.Print("Masukkan kelas : ")
	if _, err := fmt.Scan(&siswa.Kelas); err != nil || siswa.Kelas == 99 {
		utils.ClearScreen()
		Home()
		return
	}

	fmt.Print("Masukkan jurusan : ")
	if _, err := fmt.Scan(&siswa.Jurusan); err != nil || siswa.Jurusan == "99" {
		utils.ClearScreen()
		Home()
		return
	}

	model.DataSiswa = append(model.DataSiswa, siswa)
	fmt.Println("Data siswa berhasil ditambahkan!")

	if err := utils.SaveSiswa("Model/data_siswa.json"); err != nil {
		fmt.Println("Error saving siswa data:", err)
	}
}
