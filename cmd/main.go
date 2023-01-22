package main

import (
	"os"

	config "github.com/fakriardian/staffinc/internal/database"
	"github.com/fakriardian/staffinc/internal/delivery/rest"
	hargaRepository "github.com/fakriardian/staffinc/internal/repository/harga"
	rekeningRepository "github.com/fakriardian/staffinc/internal/repository/rekening"
	eUseCase "github.com/fakriardian/staffinc/internal/use-case/emas"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database := config.GetDb(os.Getenv("DB_ADDRESS"))

	hargaRepo := hargaRepository.GetRepository(database)
	rekeningRepo := rekeningRepository.GetRepository(database)

	emasUseCase := eUseCase.GetUseCase(hargaRepo, rekeningRepo)
	go emasUseCase.ConsumerUpdateHarga()

	handler := rest.NewHandler(emasUseCase)

	// rest.LoadMiddleware(e)
	rest.LoadRoutes(e, handler)

	e.Logger.Fatal(e.Start(":5000"))
}
