package handler

import (
	"github.com/labstack/echo"
	"julo_test/service"
	"net/http"
)

type AccountHandler struct {
	e              *echo.Echo
	accountUsecase service.AccountUsecase
}

func NewAccountHandler(
	e *echo.Echo,
	accountUsecase service.AccountUsecase,
) *AccountHandler {
	return &AccountHandler{
		e:              e,
		accountUsecase: accountUsecase,
	}
}

func (h *AccountHandler) Create(e echo.Context) error {
	return h.accountUsecase.Create()
}

func (h *AccountHandler) FindByID(e echo.Context) error {
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

	res, err := h.accountUsecase.FindByID(id)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, res)
}
