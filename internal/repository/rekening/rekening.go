package rekening

import (
	"fmt"
	"math"

	"github.com/fakriardian/staffinc/internal/model"
	"github.com/fakriardian/staffinc/internal/model/constant"
	"gorm.io/gorm"
)

type rekeningRepo struct {
	db *gorm.DB
}

func GetRepository(
	db *gorm.DB,
) Repository {
	return &rekeningRepo{
		db: db,
	}
}

func (rr *rekeningRepo) CheckSaldo(norek string) (model.Rekening, error) {
	var rekeningData model.Rekening

	if err := rr.db.Where(model.Rekening{NoRek: norek}).First(&rekeningData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Rekening{}, nil
		} else {
			return model.Rekening{}, err
		}
	}

	return rekeningData, nil
}

func (rr *rekeningRepo) IsExisting(norek string) (bool, error) {
	var rekeningData model.Rekening

	if err := rr.db.Where(model.Rekening{NoRek: norek}).First(&rekeningData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, fmt.Errorf("norek: %s not available", norek)
		} else {
			return false, err
		}
	}

	return rekeningData.NoRek != "", nil
}

func (rr *rekeningRepo) CreateRek(dataRek model.Rekening) (model.Rekening, error) {
	if err := rr.db.Create(&dataRek).Error; err != nil {
		return model.Rekening{}, err
	}

	return dataRek, nil
}

func (rr *rekeningRepo) UpdateSaldo(norek, tipe string, saldo float64) (model.Rekening, error) {
	var rekeningData model.Rekening
	var totalSaldo float64

	existingRekeningData, err := rr.CheckSaldo(norek)
	if err != nil {
		return model.Rekening{}, err
	}

	if existingRekeningData.NoRek == "" {
		rr.CreateRek(model.Rekening{
			NoRek: norek,
			Saldo: saldo,
		})
	}

	if tipe == constant.TransactionTypeTopUp {
		totalSaldo = math.Floor((existingRekeningData.Saldo+saldo)*1000) / 1000
	} else {
		totalSaldo = math.Floor((existingRekeningData.Saldo-saldo)*1000) / 1000
	}

	if err := rr.db.Where(model.Rekening{NoRek: norek}).Updates(model.Rekening{
		Saldo: totalSaldo,
	}).Find(&rekeningData).Error; err != nil {
		return model.Rekening{}, err
	}

	return rekeningData, nil

}
