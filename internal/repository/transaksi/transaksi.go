package transaksi

import (
	"fmt"

	"github.com/fakriardian/staffinc/internal/model"
	"gorm.io/gorm"
)

type transactionRepo struct {
	db *gorm.DB
}

func GetRepository(
	db *gorm.DB,
) Repository {
	return &transactionRepo{
		db: db,
	}
}

func (tr *transactionRepo) AddTrasaction(data model.Transaction) (model.Transaction, error) {
	if err := tr.db.Create(&data).Error; err != nil {
		return model.Transaction{}, err
	}

	return data, nil
}

func (tr *transactionRepo) GetTransaction(norek string, startDate, endDate int64) ([]model.Transaction, error) {
	var transactionData []model.Transaction

	if startDate != 0 && endDate != 0 {
		if err := tr.db.Where(
			"no_rek = ? AND transaction_date BETWEEN ? AND ?", norek, startDate, endDate,
		).Find(&transactionData).Error; err != nil {
			return nil, err
		}
	} else if startDate != 0 {
		if err := tr.db.Where(
			"no_rek = ? AND transaction_date > ?", norek, startDate,
		).Find(&transactionData).Error; err != nil {
			return nil, err
		}
	} else if endDate != 0 {
		if err := tr.db.Where(
			"no_rek = ? AND transaction_date < ?", norek, endDate,
		).Find(&transactionData).Error; err != nil {
			return nil, err
		}
	} else {
		if err := tr.db.Where(model.Transaction{NoRek: norek}).Find(&transactionData).Error; err != nil {
			return nil, err
		}
	}

	return transactionData, nil
}

func (tr *transactionRepo) IsExisting(norek string) (bool, error) {
	var transactionData model.Transaction

	if err := tr.db.Where(model.Transaction{NoRek: norek}).First(&transactionData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, fmt.Errorf("norek: %s not available", norek)
		} else {
			return false, err
		}
	}

	return transactionData.NoRek != "", nil
}
