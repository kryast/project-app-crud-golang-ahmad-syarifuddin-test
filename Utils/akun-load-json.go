package utils

import (
	"encoding/json"
	"os"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
)

func LoadAkun(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {

			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&model.DataAkun); err != nil {
		return err
	}

	return nil
}
