package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func CreateGuru() {
	var guru model.Guru
	guru.ID = utils.TambahID(model.DataGuru)
	fmt.Print("Masukkan nama guru: ")
	fmt.Scan(&guru.Nama)
	fmt.Print("Masukkan NIP guru: ")
	fmt.Scan(&guru.NIP)
	fmt.Print("Masukkan mata pelajaran: ")
	fmt.Scan(&guru.MataPelajaran)
	model.DataGuru = append(model.DataGuru, guru)
	fmt.Println("Data guru berhasil ditambahkan!")

	if err := utils.SaveGuru("Model/data_guru.json"); err != nil {
		fmt.Println("Error saving guru data:", err)
	}
}
