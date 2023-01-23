package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fakriardian/staffinc/internal/model/constant"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *handler) CheckMutasi(c echo.Context) error {
	var request constant.CheckMutasiRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	CheckMutasi, err := h.emasUseCase.CheckMutasi(request)
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
		"data":  CheckMutasi,
	})

}

func (h *handler) ProducerTopUp(c echo.Context) error {
	var request constant.TopUpRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	dataTopUp, err := h.emasUseCase.ProducerTopUp(request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
			"reff_id": uuid.NewString(),
		})
	}

	fmt.Printf("added new record by norek: %s\n", dataTopUp.NoRek)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"reff_id": uuid.NewString(),
	})
}

func (h *handler) ProducerBuyBack(c echo.Context) error {
	var request constant.BuyBackRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	dataBuyBack, err := h.emasUseCase.ProducerBuyBack(request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
			"reff_id": uuid.NewString(),
		})
	}

	fmt.Printf("added new record by norek: %s\n", dataBuyBack.NoRek)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"reff_id": uuid.NewString(),
	})
}
