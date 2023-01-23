package model

import (
	"github.com/fakriardian/staffinc/internal/model/constant"
)

type Rekening struct {
	NoRek string  `gorm:"primaryKey unique" json:"norek"`
	Saldo float64 `json:"saldo"`
}

type Transaction struct {
	ID              string                   `gorm:"primaryKey" json:"-"`
	NoRek           string                   `json:"norek"`
	TransactionDate int64                    `json:"date"`
	Type            constant.TransactionType `json:"type"`
	Gram            float64                  `json:"gram"`
	Saldo           float64                  `json:"saldo"`
	HargaTopUp      int64                    `json:"harga_topup"`
	HargaBuyBack    int64                    `json:"harga_buyback"`
}
