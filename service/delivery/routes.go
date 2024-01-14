package delivery

import (
	"github.com/labstack/echo"
	"julo_test/service"
	"julo_test/service/delivery/handler"
)

func NewRouter(
	e *echo.Echo,
	walletUsecase service.WalletUsecase,
	transactionUsecase service.TransactionUsecase,
) {
	w := handler.NewWalletHandler(e, walletUsecase)
	trx := handler.NewTransactionHandler(e, transactionUsecase, walletUsecase)

	rr := e.Group("/api/v1")
	wc := rr.Group("/wallet")

	rr.POST("/init", w.InitWallet)
	wc.POST("", w.EnableWallet)
	wc.PATCH("", w.DisableWallet)
	wc.GET("", w.ViewWallet)
	wc.POST("/deposits", w.Deposit)
	wc.POST("/withdrawals", w.Withdraw)

	wc.GET("/transactions", trx.ViewWalletTransaction)
}
