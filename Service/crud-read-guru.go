package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func ReadGuru() {
	if err := utils.LoadGuru("Model/data_guru.json"); err != nil {
		fmt.Println("Error loading guru data:", err)
	}

	fmt.Println("Data Guru:")
	for _, guru := range model.DataGuru {
		fmt.Printf("ID: %d, Nama: %s, NIP: %d, Mata Pelajaran: %s\n", guru.ID, guru.Nama, guru.NIP, guru.MataPelajaran)
	}

}
