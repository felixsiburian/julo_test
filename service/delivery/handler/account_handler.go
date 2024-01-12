package handler

import (
	"github.com/labstack/echo"
	"julo_test/service"
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
