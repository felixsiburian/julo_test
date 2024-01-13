package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"julo_test/service"
	"net/http"
)

type WalletHandler struct {
	e             *echo.Echo
	walletUsecase service.WalletUsecase
}

func NewWalletHandler(
	e *echo.Echo,
	walletUsecase service.WalletUsecase,
) *WalletHandler {
	return &WalletHandler{
		e:             e,
		walletUsecase: walletUsecase,
	}
}

func (h *WalletHandler) InitWallet(e echo.Context) error {
	id := e.FormValue("customer_xid")
	if id == "" {
		resMap := make(map[string]interface{})
		errorMap := make(map[string]interface{})
		res := make(map[string]interface{})

		resMap["customer_xid"] = []string{"Missing data for required field."}
		errorMap["error"] = resMap
		res["data"] = errorMap
		res["status"] = "fail"
		return e.JSON(http.StatusBadRequest, res)
	}

	parseID := uuid.MustParse(id)
	res, err := h.walletUsecase.InitWallet(parseID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, res)
}
