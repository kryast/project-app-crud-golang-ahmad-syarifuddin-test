package view

import (
	"errors"
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func generateID(data interface{}) int {
	var id int
	switch d := data.(type) {
	case []model.Guru:
		id = len(d) + 1
	case []model.Siswa:
		id = len(d) + 1
	}
	return id
}

func Crud() {
	defer func() {
		if r := recover(); r != nil {
			utils.ClearScreen()
			fmt.Println("Terjadi kesalahan yang tidak terduga")
			Crud()
		}
	}()

	var input int
	utils.CrudMenu()
	fmt.Scan(&input)

	switch input {
	case 1: // Input Data Guru
		var guru model.Guru
		guru.ID = generateID(model.DataGuru)
		fmt.Print("Masukkan nama guru: ")
		fmt.Scan(&guru.Nama)
		fmt.Print("Masukkan NIP guru: ")
		fmt.Scan(&guru.NIP)
		fmt.Print("Masukkan mata pelajaran: ")
		fmt.Scan(&guru.MataPelajaran)
		model.DataGuru = append(model.DataGuru, guru)
		fmt.Println("Data guru berhasil ditambahkan!")

	case 2: // Input Data Siswa
		var siswa model.Siswa
		siswa.ID = generateID(model.DataSiswa)
		fmt.Print("Masukkan nama siswa: ")
		fmt.Scan(&siswa.Nama)
		fmt.Print("Masukkan NIS siswa: ")
		fmt.Scan(&siswa.NIS)
		fmt.Print("Masukkan kelas: ")
		fmt.Scan(&siswa.Kelas)
		fmt.Print("Masukkan jurusan: ")
		fmt.Scan(&siswa.Jurusan)
		model.DataSiswa = append(model.DataSiswa, siswa)
		fmt.Println("Data siswa berhasil ditambahkan!")

	case 3: // Update Data Guru
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
				return
			}
		}
		fmt.Println("ID tidak valid!")

	case 4: // Update Data Siswa
		var id int
		fmt.Print("Masukkan ID siswa yang ingin diupdate: ")
		fmt.Scan(&id)
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
				return
			}
		}
		fmt.Println("ID tidak valid!")

	case 5: // Delete Data Guru
		var id int
		fmt.Print("Masukkan ID guru yang ingin dihapus: ")
		fmt.Scan(&id)
		for i, guru := range model.DataGuru {
			if guru.ID == id {
				model.DataGuru = append(model.DataGuru[:i], model.DataGuru[i+1:]...)
				fmt.Println("Data guru berhasil dihapus!")
				return
			}
		}
		fmt.Println("ID tidak valid!")

	case 6: // Delete Data Siswa
		var id int
		fmt.Print("Masukkan ID siswa yang ingin dihapus: ")
		fmt.Scan(&id)
		for i, siswa := range model.DataSiswa {
			if siswa.ID == id {
				model.DataSiswa = append(model.DataSiswa[:i], model.DataSiswa[i+1:]...)
				fmt.Println("Data siswa berhasil dihapus!")
				return
			}
		}
		fmt.Println("ID tidak valid!")

	case 7: // Tampilkan Semua Data Guru
		fmt.Println("Data Guru:")
		for _, guru := range model.DataGuru {
			fmt.Printf("ID: %d, Nama: %s, NIP: %d, Mata Pelajaran: %s\n", guru.ID, guru.Nama, guru.NIP, guru.MataPelajaran)
		}

	case 8: // Tampilkan Semua Data Siswa
		fmt.Println("Data Siswa:")
		for _, siswa := range model.DataSiswa {
			fmt.Printf("ID: %d, Nama: %s, NIS: %d, Kelas: %d, Jurusan: %s\n", siswa.ID, siswa.Nama, siswa.NIS, siswa.Kelas, siswa.Jurusan)
		}

	case 0:
		utils.ClearScreen()
		Home()

	case 99:
		utils.ClearScreen()
		fmt.Println("Terima kasih! Sampai jumpa!")
		return // Keluar dari program

	default:
		err := errors.New("pilihan tidak valid")
		utils.Panik(err)
	}

	Crud() // Tampilkan menu lagi setelah selesai
}
