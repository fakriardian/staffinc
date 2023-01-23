package emas

import (
	"time"

	"github.com/fakriardian/staffinc/internal/model"
	"github.com/fakriardian/staffinc/internal/model/constant"
	"github.com/google/uuid"
)

func (eu *emasUseCase) ConsumerUpdateHarga(request constant.InputHargaRequest) (model.Harga, error) {
	isExistingData, err := eu.hargaRepo.IsExisting()
	if err != nil {
		return model.Harga{}, err
	}

	if isExistingData != "" {
		eu.hargaRepo.DeleteExisting(isExistingData)
	}

	hargaData, err := eu.hargaRepo.AddHarga(model.Harga{
		AdminID:      request.AdminID,
		HargaTopUp:   request.HargaTopUp,
		HargaBuyBack: request.HargaBuyBack,
	})

	if err != nil {
		return model.Harga{}, err
	}

	return hargaData, nil
}

func (eu *emasUseCase) ConsumerTopUp(request constant.TopUpRequest) (model.Transaction, error) {
	validationHarga, err := eu.hargaRepo.CheckHarga()
	if err != nil {
		return model.Transaction{}, err
	}

	updateSaldo, err := eu.rekeningRepo.UpdateSaldo(request.NoRek, constant.TransactionTypeTopUp, request.Gram)
	if err != nil {
		return model.Transaction{}, err
	}

	transactionData, err := eu.transactionRepo.AddTrasaction(model.Transaction{
		ID:              uuid.NewString(),
		NoRek:           request.NoRek,
		TransactionDate: time.Now().Unix(),
		Type:            constant.TransactionTypeTopUp,
		Gram:            request.Gram,
		Saldo:           updateSaldo.Saldo,
		HargaTopUp:      request.HargaTopUp,
		HargaBuyBack:    validationHarga.HargaBuyBack,
	})

	if err != nil {
		return model.Transaction{}, err
	}

	return transactionData, nil
}

func (eu *emasUseCase) ConsumerBuyBack(request constant.BuyBackRequest) (model.Transaction, error) {
	validationHarga, err := eu.hargaRepo.CheckHarga()
	if err != nil {
		return model.Transaction{}, err
	}

	updateSaldo, err := eu.rekeningRepo.UpdateSaldo(request.NoRek, constant.TransactionTypeBuyBack, request.Gram)

	if err != nil {
		return model.Transaction{}, err
	}

	transactionData, err := eu.transactionRepo.AddTrasaction(model.Transaction{
		ID:              uuid.NewString(),
		NoRek:           request.NoRek,
		TransactionDate: time.Now().Unix(),
		Type:            constant.TransactionTypeBuyBack,
		Gram:            request.Gram,
		Saldo:           updateSaldo.Saldo,
		HargaTopUp:      validationHarga.HargaTopUp,
		HargaBuyBack:    request.HargaBuyBack,
	})

	if err != nil {
		return model.Transaction{}, err
	}

	return transactionData, nil
}
