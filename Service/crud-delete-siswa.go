package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func DeleteSiswa() {
	var id int
	fmt.Print("Masukkan ID siswa yang ingin dihapus: ")
	fmt.Scan(&id)
	for i, siswa := range model.DataSiswa {
		if siswa.ID == id {
			model.DataSiswa = append(model.DataSiswa[:i], model.DataSiswa[i+1:]...)
			fmt.Println("Data siswa berhasil dihapus!")

			if err := utils.SaveSiswa("Model/data_siswa.json"); err != nil {
				fmt.Println("Error saving siswa data:", err)
			}
			return
		}
	}
	fmt.Println("ID tidak valid!")
}
