package response

type (
	FailedEnableWallet struct {
		Status string
		Data   DataFailedEnableWallet `json:"data"`
	}

	DataFailedEnableWallet struct {
		Error string `json:"error"`
	}
)
