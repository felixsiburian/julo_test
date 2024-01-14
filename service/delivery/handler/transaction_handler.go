package handler

import (
	"github.com/labstack/echo"
	"julo_test/service"
	"julo_test/service/model/response"
	"julo_test/service/tools"
	"net/http"
)

type TransactionHandler struct {
	e                  *echo.Echo
	transactionUsecase service.TransactionUsecase
	walletUsecase      service.WalletUsecase
}

func NewTransactionHandler(
	e *echo.Echo,
	transactionUsecase service.TransactionUsecase,
	walletUsecase service.WalletUsecase,
) *TransactionHandler {
	return &TransactionHandler{
		e:                  e,
		transactionUsecase: transactionUsecase,
		walletUsecase:      walletUsecase,
	}
}

func (h *TransactionHandler) ViewWalletTransaction(e echo.Context) error {
	tokenHeader := e.Request().Header.Get("Authorization")
	if tokenHeader == "" {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": "invalid token",
		})
	}

	token, err := tools.PartitionToken(tokenHeader)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": err,
		})
	}

	walletId, err := tools.Decrypt(token)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": err,
		})
	}

	walletData, err := h.walletUsecase.FindWalletByWalletID(walletId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": err,
		})
	}

	if walletData.Data.Wallet.Status != "enabled" {
		errorsData := response.FailedEnableWallet{
			Status: "failed",
			Data: response.DataFailedEnableWallet{
				Error: "Wallet disabled",
			},
		}
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": errorsData,
		})
	}

	res, err := h.transactionUsecase.FindTransactionByWalletId(walletId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": err,
		})
	}
	res.Status = "success"

	return e.JSON(http.StatusOK, res)
}
