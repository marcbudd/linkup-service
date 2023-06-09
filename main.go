package main

import (
	_ "github.com/marcbudd/linkup-service/docs"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/router"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDb()
	initalizers.SyncDatabase()
}

func main() {

	r := router.SetupRouter()
	r.Run()
	defer initalizers.CloseDbConnection()

}
