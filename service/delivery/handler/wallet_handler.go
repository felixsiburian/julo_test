package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"julo_test/service"
	"julo_test/service/model/request"
	"julo_test/service/model/response"
	"julo_test/service/tools"
	"net/http"
	"strconv"
	"strings"
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

func (h *WalletHandler) EnableWallet(e echo.Context) error {
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

	if strings.ToLower(walletData.Data.Wallet.Status) != "enabled" {
		res, err := h.walletUsecase.EnableWallet(walletId)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status":  "failed",
				"message": err,
			})
		}

		return e.JSON(http.StatusOK, res)
	} else {
		res := response.FailedEnableWallet{
			Status: "failed",
			Data: response.DataFailedEnableWallet{
				Error: "Already enabled",
			},
		}
		return e.JSON(http.StatusBadRequest, res)
	}
}

func (h *WalletHandler) ViewWallet(e echo.Context) error {
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
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
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
	walletData.Status = "success"

	return e.JSON(http.StatusOK, walletData)
}

func (h *WalletHandler) Deposit(e echo.Context) error {
	tokenHeader := e.Request().Header.Get("Authorization")
	if tokenHeader == "" {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": "invalid token",
		})
	}

	refId := e.FormValue("reference_id")
	if refId == "" {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": "invalid refId",
		})
	}

	amount := e.FormValue("amount")
	if amount == "" {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": "invalid amount",
		})
	}

	amountFloat, _ := strconv.ParseFloat(amount, 10)

	token, err := tools.PartitionToken(tokenHeader)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": err,
		})
	}

	walletId, err := tools.Decrypt(token)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
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

	convertWalletId := uuid.MustParse(walletId)
	convertRefId := uuid.MustParse(refId)

	finalAmount := walletData.Data.Wallet.Balance + amountFloat

	res, err := h.walletUsecase.UpdateBalance(request.UpdateWalletRequest{
		WalletId:    convertWalletId,
		Amount:      finalAmount,
		AmountAdded: amountFloat,
		Type:        "deposit",
		RefId:       convertRefId,
		DepositBy:   walletData.Data.Wallet.OwnedBy,
	})
	res.Deposit.Amount = amountFloat

	return e.JSON(http.StatusOK, res)
}

func (h *WalletHandler) Withdraw(e echo.Context) error {
	tokenHeader := e.Request().Header.Get("Authorization")
	if tokenHeader == "" {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": "invalid token",
		})
	}

	refId := e.FormValue("reference_id")
	if refId == "" {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": "invalid refId",
		})
	}

	amount := e.FormValue("amount")
	if amount == "" {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": "invalid amount",
		})
	}

	amountFloat, _ := strconv.ParseFloat(amount, 10)

	token, err := tools.PartitionToken(tokenHeader)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": err,
		})
	}

	walletId, err := tools.Decrypt(token)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
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

	convertWalletId := uuid.MustParse(walletId)
	convertRefId := uuid.MustParse(refId)

	if walletData.Data.Wallet.Balance < amountFloat {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "invalid balance",
		})
	}

	finalAmount := walletData.Data.Wallet.Balance - amountFloat

	res, err := h.walletUsecase.UpdateBalance(request.UpdateWalletRequest{
		WalletId:    convertWalletId,
		Amount:      finalAmount,
		AmountAdded: amountFloat,
		Type:        "withdraw",
		RefId:       convertRefId,
		DepositBy:   walletData.Data.Wallet.OwnedBy,
	})
	res.Deposit.Amount = amountFloat

	return e.JSON(http.StatusOK, res)
}
