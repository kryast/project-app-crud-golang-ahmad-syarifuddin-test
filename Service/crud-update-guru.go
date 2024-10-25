package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func UpdateGuru() {
	var id int
	fmt.Print("Masukkan ID guru yang ingin diupdate: ")
	fmt.Scan(&id)
	for i, guru := range model.DataGuru {
		if guru.ID == id {
			fmt.Print("Masukkan nama guru baru: ")
			fmt.Scan(&model.DataGuru[i].Nama)
			fmt.Print("Masukkan NIP guru baru: ")
			fmt.Scan(&model.DataGuru[i].NIP)
			fmt.Print("Masukkan mata pelajaran baru: ")
			fmt.Scan(&model.DataGuru[i].MataPelajaran)
			fmt.Println("Data guru berhasil diupdate!")

			if err := utils.SaveGuru("Model/data_guru.json"); err != nil {
				fmt.Println("Error saving guru data:", err)
			}
			return
		}
	}
	fmt.Println("ID tidak valid!")
}
