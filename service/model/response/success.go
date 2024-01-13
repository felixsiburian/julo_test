package response

import "julo_test/service/model"

type (
	SuccessInitWallet struct {
		Data   DataSuccessInitWallet `json:"data"`
		Status string                `json:"status"`
	}

	DataSuccessInitWallet struct {
		Token string `json:"token"`
	}

	SuccessEnableWallet struct {
		Status string                  `json:"status"`
		Data   DataSuccessEnableWallet `json:"data"`
	}

	DataSuccessEnableWallet struct {
		Wallet model.Wallet `json:"wallet"`
	}
)
