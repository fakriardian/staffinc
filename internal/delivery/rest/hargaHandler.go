package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fakriardian/staffinc/internal/model/constant"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *handler) ProducerInputHarga(c echo.Context) error {
	var request constant.InputHargaRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	hargaData, err := h.emasUseCase.ProducerUpdateHarga(request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
			"reff_id": uuid.NewString(),
		})
	}

	fmt.Printf("added new record by adminId: %s\n", hargaData.AdminID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"reff_id": uuid.NewString(),
	})
}

func (h *handler) CheckHarga(c echo.Context) error {
	checkHarga, err := h.emasUseCase.CheckHarga()
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
		"data":  checkHarga,
	})
}
