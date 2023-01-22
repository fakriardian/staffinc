package database

import (
	"github.com/fakriardian/staffinc/internal/model"
	"gorm.io/gorm"
)

func migrationDB(db *gorm.DB) {

	// for migration
	db.AutoMigrate(&model.Harga{}, &model.Rekening{}, &model.Transaction{})

}
