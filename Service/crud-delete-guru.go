package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func DeleteGuru() {
	var id int
	fmt.Print("Masukkan ID guru yang ingin dihapus: ")
	fmt.Scan(&id)
	for i, guru := range model.DataGuru {
		if guru.ID == id {
			model.DataGuru = append(model.DataGuru[:i], model.DataGuru[i+1:]...)
			fmt.Println("Data guru berhasil dihapus!")

			if err := utils.SaveGuru("Model/data_guru.json"); err != nil {
				fmt.Println("Error saving guru data:", err)
			}
			return
		}
	}
	fmt.Println("ID tidak valid!")
}
