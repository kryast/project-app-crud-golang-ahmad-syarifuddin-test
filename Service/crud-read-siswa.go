package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func ReadSiswa() {
	if err := utils.LoadSiswa("Model/data_siswa.json"); err != nil {
		fmt.Println("Error loading siswa data:", err)
	}

	fmt.Println("Data Siswa:")
	for _, siswa := range model.DataSiswa {
		fmt.Printf("ID: %d, Nama: %s, NIS: %d, Kelas: %d, Jurusan: %s\n", siswa.ID, siswa.Nama, siswa.NIS, siswa.Kelas, siswa.Jurusan)
	}
}
