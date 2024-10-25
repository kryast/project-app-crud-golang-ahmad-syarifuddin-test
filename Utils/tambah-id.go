package utils

import model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"

func TambahID(data interface{}) int {
	var id int
	switch d := data.(type) {
	case []model.Guru:
		id = len(d) + 1
	case []model.Siswa:
		id = len(d) + 1
	}
	return id
}
