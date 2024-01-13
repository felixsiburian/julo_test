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
