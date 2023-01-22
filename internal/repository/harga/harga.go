package harga

import (
	"github.com/fakriardian/staffinc/internal/model"
	"gorm.io/gorm"
)

type hargaRepo struct {
	db *gorm.DB
}

func GetRepository(
	db *gorm.DB,
) Repository {
	return &hargaRepo{
		db: db,
	}
}

func (hr *hargaRepo) AddHarga(hargaData model.Harga) (model.Harga, error) {
	if err := hr.db.Create(&hargaData).Error; err != nil {
		return model.Harga{}, err
	}

	return hargaData, nil
}

func (hr *hargaRepo) IsExisting() (string, error) {
	var hargaData model.Harga

	if err := hr.db.First(&hargaData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		} else {
			return "", err
		}
	}

	return hargaData.AdminID, nil
}

func (hr *hargaRepo) DeleteExisting(adminID string) (string, error) {

	if err := hr.db.Unscoped().Where(model.Harga{AdminID: adminID}).Delete(model.Harga{}).Error; err != nil {

		return "failed", err
	}

	return "done", nil

}
