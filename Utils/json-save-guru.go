package utils

import (
	"encoding/json"
	"os"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
)

func SaveGuru(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	return encoder.Encode(model.DataGuru)
}
