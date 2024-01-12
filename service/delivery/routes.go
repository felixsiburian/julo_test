package delivery

import (
	"github.com/labstack/echo"
	"julo_test/service"
	"julo_test/service/delivery/handler"
)

func NewRouter(
	e *echo.Echo,
	accountUsecase service.AccountUsecase,
) {
	a := handler.NewAccountHandler(e, accountUsecase)

	rr := e.Group("/api/v1")
	ac := rr.Group("/account")

	ac.POST("", a.Create)
}
