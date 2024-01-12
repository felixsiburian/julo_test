package service

type IAccountRepository interface {
	Create() error
}

type ITransactionRepository interface {
}

type IWalletRepository interface {
}
