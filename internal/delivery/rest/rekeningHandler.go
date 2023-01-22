package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fakriardian/staffinc/internal/model/constant"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *handler) CheckSaldo(c echo.Context) error {
	var request constant.CheckSaldoRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	checkRekening, err := h.emasUseCase.CheckRekening(request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
			"reff_id": uuid.NewString(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error": false,
		"data":  checkRekening,
	})

}
