package transaksi

import "github.com/fakriardian/staffinc/internal/model"

type Repository interface {
	AddTrasaction(data model.Transaction) (model.Transaction, error)
	GetTransaction(norek string, startData, endData int64) ([]model.Transaction, error)
	IsExisting(norek string) (bool, error)
}
