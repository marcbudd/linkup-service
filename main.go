package main

import (
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

}
