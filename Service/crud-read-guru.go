package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
)

func ReadGuru() {
	fmt.Println("Data Guru:")
	for _, guru := range model.DataGuru {
		fmt.Printf("ID: %d, Nama: %s, NIP: %d, Mata Pelajaran: %s\n", guru.ID, guru.Nama, guru.NIP, guru.MataPelajaran)
	}
}
