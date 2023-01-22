package rekening

import (
	"github.com/fakriardian/staffinc/internal/model"
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
