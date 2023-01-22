package rest

import "github.com/fakriardian/staffinc/internal/use-case/emas"

type handler struct {
	emasUseCase emas.Usecase
}

func NewHandler(emasUseCase emas.Usecase) *handler {
	return &handler{
		emasUseCase,
	}
}
