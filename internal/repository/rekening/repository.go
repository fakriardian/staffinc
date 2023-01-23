package rekening

import "github.com/fakriardian/staffinc/internal/model"

type Repository interface {
	CheckSaldo(norek string) (model.Rekening, error)
	IsExisting(norek string) (bool, error)
	CreateRek(dataRek model.Rekening) (model.Rekening, error)
	UpdateSaldo(norek, tipe string, saldo float64) (model.Rekening, error)
}
