package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func CreateSiswa() {
	var siswa model.Siswa
	siswa.ID = utils.TambahID(model.DataSiswa)
	fmt.Print("Masukkan nama siswa: ")
	fmt.Scan(&siswa.Nama)
	fmt.Print("Masukkan NIS siswa: ")
	fmt.Scan(&siswa.NIS)
	fmt.Print("Masukkan kelas: ")
	fmt.Scan(&siswa.Kelas)
	fmt.Print("Masukkan jurusan: ")
	fmt.Scan(&siswa.Jurusan)
	model.DataSiswa = append(model.DataSiswa, siswa)
	fmt.Println("Data siswa berhasil ditambahkan!")

	if err := utils.SaveSiswa("Model/data_siswa.json"); err != nil {
		fmt.Println("Error saving siswa data:", err)
	}
}
