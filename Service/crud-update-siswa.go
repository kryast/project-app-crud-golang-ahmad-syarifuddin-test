package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func UpdateSiswa() {
	ReadSiswa()
	fmt.Println("\n0. Kembali")
	var id int
	fmt.Print("Masukkan ID siswa yang ingin diupdate: ")
	fmt.Scan(&id)
	if id == 0 {
		utils.ClearScreen()
		Crud()
	} else if id == 99 {
		utils.ClearScreen()
		Home()
	}
	for i, siswa := range model.DataSiswa {
		if siswa.ID == id {
			fmt.Print("Masukkan nama siswa baru: ")
			fmt.Scan(&model.DataSiswa[i].Nama)
			fmt.Print("Masukkan NIS siswa baru: ")
			fmt.Scan(&model.DataSiswa[i].NIS)
			fmt.Print("Masukkan kelas baru: ")
			fmt.Scan(&model.DataSiswa[i].Kelas)
			fmt.Print("Masukkan jurusan baru: ")
			fmt.Scan(&model.DataSiswa[i].Jurusan)
			fmt.Println("Data siswa berhasil diupdate!")

			if err := utils.SaveSiswa("Model/data_siswa.json"); err != nil {
				fmt.Println("Error saving siswa data:", err)
			}
			return
		}
	}
	fmt.Println("ID tidak valid!")
}
