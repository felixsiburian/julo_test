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
	w := handler.NewWalletHandler(e, walletUsecase)

	rr := e.Group("/api/v1")
	wc := rr.Group("/wallet")

	rr.POST("/init", w.InitWallet)
	wc.POST("", w.EnableWallet)
	wc.GET("", w.ViewWallet)
}
