package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
	view "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/View"
)

func Crud() {
	defer func() {
		if r := recover(); r != nil {
			utils.ClearScreen()
			fmt.Println("Terjadi kesalahan yang tidak terduga")
			Crud()
		}
	}()

	if err := utils.LoadGuru("Model/data_guru.json"); err != nil {
		fmt.Println("Error loading guru data:", err)
	}
	if err := utils.LoadSiswa("Model/data_siswa.json"); err != nil {
		fmt.Println("Error loading siswa data:", err)
	}

	parentCtx := context.Background()
	deadline := time.Now().Add(20 * time.Second)

	ctx, cancel := context.WithDeadline(parentCtx, deadline)
	defer cancel()

	go func() {
		<-ctx.Done()
		utils.ClearScreen()
		fmt.Println("99. Kembali")
		fmt.Println("Sesi anda habis. Silahkan Login Kembali")
		fmt.Print("Masukkan angka :")
	}()
	var input int
	view.CrudMenu()
	fmt.Scan(&input)

	switch input {

	case 1:
		utils.ClearScreen()
		CreateGuru()
		BackCrud()

	case 2:
		utils.ClearScreen()
		CreateSiswa()
		BackCrud()

	case 3:
		utils.ClearScreen()
		UpdateGuru()
		BackCrud()

	case 4:
		utils.ClearScreen()
		UpdateSiswa()
		BackCrud()

	case 5:
		utils.ClearScreen()
		DeleteGuru()
		BackCrud()

	case 6:
		utils.ClearScreen()
		DeleteSiswa()
		BackCrud()

	case 7:
		utils.ClearScreen()
		ReadGuru()
		BackCrud()

	case 8:
		utils.ClearScreen()
		ReadSiswa()
		BackCrud()

	case 0:
		utils.ClearScreen()
		Home()

	case 99:
		utils.ClearScreen()
		Home()

	default:
		err := errors.New("pilihan tidak valid")
		utils.Panik(err)
	}

}
