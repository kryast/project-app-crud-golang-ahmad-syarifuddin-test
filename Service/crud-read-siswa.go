package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
)

func ReadSiswa() {
	fmt.Println("Data Siswa:")
	for _, siswa := range model.DataSiswa {
		fmt.Printf("ID: %d, Nama: %s, NIS: %d, Kelas: %d, Jurusan: %s\n", siswa.ID, siswa.Nama, siswa.NIS, siswa.Kelas, siswa.Jurusan)
	}
}
