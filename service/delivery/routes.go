package delivery

import (
	"github.com/labstack/echo"
	"julo_test/service"
	"julo_test/service/delivery/handler"
)

func NewRouter(
	e *echo.Echo,
	accountUsecase service.AccountUsecase,
	walletUsecase service.WalletUsecase,
) {
	a := handler.NewAccountHandler(e, accountUsecase)
	w := handler.NewWalletHandler(e, walletUsecase)

	rr := e.Group("/api/v1")
	ac := rr.Group("/account")
	//wc := rr.Group("/wallet")

	ac.POST("", a.Create)
	ac.GET("/find", a.FindByID)

	rr.POST("/init", w.InitWallet)

}
