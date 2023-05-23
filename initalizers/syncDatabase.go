package initalizers

import "github.com/marcbudd/linkup-service/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Tweet{})
}
