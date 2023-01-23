package main

import (
	"os"

	config "github.com/fakriardian/staffinc/internal/database"
	con "github.com/fakriardian/staffinc/internal/delivery/kafka/consumer"
	"github.com/fakriardian/staffinc/internal/delivery/rest"
	hargaRepository "github.com/fakriardian/staffinc/internal/repository/harga"
	rekeningRepository "github.com/fakriardian/staffinc/internal/repository/rekening"
	transaksiRepository "github.com/fakriardian/staffinc/internal/repository/transaksi"
	eUseCase "github.com/fakriardian/staffinc/internal/use-case/emas"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database := config.GetDb(os.Getenv("DB_ADDRESS"))

	hargaRepo := hargaRepository.GetRepository(database)
	rekeningRepo := rekeningRepository.GetRepository(database)
	transactionRepo := transaksiRepository.GetRepository(database)

	emasUseCase := eUseCase.GetUseCase(hargaRepo, rekeningRepo, transactionRepo)

	handler := rest.NewHandler(emasUseCase)
	kafka := con.NewHandler(emasUseCase)

	go kafka.ConsumerUpdateHarga()
	go kafka.ConsumerTopUpSalod()
	go kafka.ConsumerBuyBack()

	// rest.LoadMiddleware(e)
	rest.LoadRoutes(e, handler)

	e.Logger.Fatal(e.Start(":5000"))
}
