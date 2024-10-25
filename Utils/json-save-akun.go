package utils

import (
	"encoding/json"
	"os"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
)

func SaveAkun(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(model.DataAkun); err != nil {
		return err
	}

	return nil
}
