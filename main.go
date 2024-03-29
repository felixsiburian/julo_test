package main

import (
	"fmt"
	"github.com/labstack/echo"
	db2 "julo_test/lib/db"
	"julo_test/service/config"
	"julo_test/service/delivery"
	"julo_test/service/repository"
	"julo_test/service/usecase"
	"log"
	"os"
)

func main() {
	start()
}

func start() {
	app := config.Config{}
	e := echo.New()

	app.CatchError(app.InitEnv())
	dbConfig := app.GetDBConfig()

	dbConn := db2.ConnectionGorm(dbConfig)

	//	register repository
	accountRepo := repository.NewServiceCenterRepository(dbConn)
	walletRepo := repository.NewWalletRepository(dbConn)
	transactionRepo := repository.NewTransactionRepository(dbConn)

	//	register usecase
	walletUsecase := usecase.NewWalletUsecase(walletRepo, accountRepo, transactionRepo)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)

	delivery.NewRouter(e, walletUsecase, transactionUsecase)

	log.Println("service running on port: ", os.Getenv("APP_PORT"))
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))))
}
