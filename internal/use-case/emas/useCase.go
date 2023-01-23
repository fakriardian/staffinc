package emas

import (
	"github.com/fakriardian/staffinc/internal/model"
	"github.com/fakriardian/staffinc/internal/model/constant"
	"github.com/fakriardian/staffinc/internal/repository/harga"
	"github.com/fakriardian/staffinc/internal/repository/rekening"
	"github.com/fakriardian/staffinc/internal/repository/transaksi"
)

type emasUseCase struct {
	hargaRepo       harga.Repository
	rekeningRepo    rekening.Repository
	transactionRepo transaksi.Repository
}

func GetUseCase(
	hargaRepo harga.Repository,
	rekeningRepo rekening.Repository,
	transactionRepo transaksi.Repository,
) Usecase {
	return &emasUseCase{
		hargaRepo:       hargaRepo,
		rekeningRepo:    rekeningRepo,
		transactionRepo: transactionRepo,
	}
}

func (eu *emasUseCase) CheckHarga() (model.Harga, error) {
	hargaData, err := eu.hargaRepo.CheckHarga()
	if err != nil {
		return model.Harga{}, err
	}

	return hargaData, nil
}

func (eu *emasUseCase) CheckRekening(request constant.CheckSaldoRequest) (model.Rekening, error) {
	_, err := eu.rekeningRepo.IsExisting(request.NoRek)
	if err != nil {
		return model.Rekening{}, err
	}

	rekeningData, err := eu.rekeningRepo.CheckSaldo(request.NoRek)
	if err != nil {
		return model.Rekening{}, err
	}

	return rekeningData, nil

}

func (eu *emasUseCase) CheckMutasi(request constant.CheckMutasiRequest) ([]model.Transaction, error) {
	_, err := eu.transactionRepo.IsExisting(request.NoRek)
	if err != nil {
		return []model.Transaction{}, err
	}

	transactionData, err := eu.transactionRepo.GetTransaction(request.NoRek, request.StartDate, request.EndDate)
	if err != nil {
		return []model.Transaction{}, err
	}

	return transactionData, nil
}
