package rekening

import "github.com/fakriardian/staffinc/internal/model"

type Repository interface {
	CheckSaldo(norek string) (model.Rekening, error)
}
