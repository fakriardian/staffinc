package emas

import (
	"github.com/fakriardian/staffinc/internal/model"
	"github.com/fakriardian/staffinc/internal/model/constant"
)

type Usecase interface {
	UpdateHarga(request constant.InputHargaRequest) (model.Harga, error)
	ConsumerUpdateHarga() error

	CheckHarga() (model.Harga, error)

	CheckRekening(request constant.CheckSaldoRequest) (model.Rekening, error)
}
