package emas

import (
	"github.com/fakriardian/staffinc/internal/model"
	"github.com/fakriardian/staffinc/internal/model/constant"
)

type Usecase interface {
	ProducerUpdateHarga(request constant.InputHargaRequest) (model.Harga, error)
	ConsumerUpdateHarga(request constant.InputHargaRequest) (model.Harga, error)

	CheckHarga() (model.Harga, error)

	CheckRekening(request constant.CheckSaldoRequest) (model.Rekening, error)

	CheckMutasi(request constant.CheckMutasiRequest) ([]model.Transaction, error)

	ProducerTopUp(request constant.TopUpRequest) (model.Transaction, error)
	ConsumerTopUp(request constant.TopUpRequest) (model.Transaction, error)

	ProducerBuyBack(request constant.BuyBackRequest) (model.Transaction, error)
	ConsumerBuyBack(request constant.BuyBackRequest) (model.Transaction, error)
}
