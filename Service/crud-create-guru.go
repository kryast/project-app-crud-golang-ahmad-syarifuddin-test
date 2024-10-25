package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func CreateGuru() {
	var guru model.Guru
	guru.ID = utils.TambahID(model.DataGuru)

	fmt.Print("Masukkan nama guru : ")
	if _, err := fmt.Scan(&guru.Nama); err != nil || guru.Nama == "99" {
		utils.ClearScreen()
		Home()
	}

	fmt.Print("Masukkan NIP guru : ")
	var nip int
	if _, err := fmt.Scan(&nip); err != nil || nip == 99 {
		utils.ClearScreen()
		Home()
		return
	}
	guru.NIP = nip

	fmt.Print("Masukkan mata pelajaran : ")
	if _, err := fmt.Scan(&guru.MataPelajaran); err != nil || guru.MataPelajaran == "99" {
		utils.ClearScreen()
		Home()
		return
	}

	model.DataGuru = append(model.DataGuru, guru)
	fmt.Println("Data guru berhasil ditambahkan!")

	if err := utils.SaveGuru("Model/data_guru.json"); err != nil {
		fmt.Println("Error saving guru data:", err)
	}
}
