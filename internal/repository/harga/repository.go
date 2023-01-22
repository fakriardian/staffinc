package harga

import "github.com/fakriardian/staffinc/internal/model"

type Repository interface {
	IsExisting() (string, error)
	AddHarga(hargaData model.Harga) (model.Harga, error)
	DeleteExisting(adminID string) (string, error)
	CheckHarga() (model.Harga, error)
}
