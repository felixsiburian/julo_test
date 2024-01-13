package response

type (
	SuccessInitWallet struct {
		Data   DataSuccessInitWallet `json:"data"`
		Status string                `json:"status"`
	}

	DataSuccessInitWallet struct {
		Token string `json:"token"`
	}
)
