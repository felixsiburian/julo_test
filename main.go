package main

import (
	"fmt"
	db2 "julo_test/lib/db"
	"julo_test/service/config"
)

func start() {
	app := config.Config{}
	//e := echo.New()

	app.CatchError(app.InitEnv())
	dbConfig := app.GetDBConfig()

	dbConn := db2.ConnectionGorm(dbConfig)
	fmt.Println("dbConn: ", dbConn.DB().Ping())
}

func main() {
	start()
}
