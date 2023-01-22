package main

import (
	"os"

	config "github.com/fakriardian/staffinc/internal/database"
	"github.com/fakriardian/staffinc/internal/delivery/rest"
	hargaRepository "github.com/fakriardian/staffinc/internal/repository/harga"
	eUseCase "github.com/fakriardian/staffinc/internal/use-case/emas"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database := config.GetDb(os.Getenv("DB_ADDRESS"))

	emasRepo := hargaRepository.GetRepository(database)

	emasUseCase := eUseCase.GetUseCase(emasRepo)
	go emasUseCase.ConsumerUpdateHarga()

	handler := rest.NewHandler(emasUseCase)

	// rest.LoadMiddleware(e)
	rest.LoadRoutes(e, handler)

	e.Logger.Fatal(e.Start(":5000"))
}
