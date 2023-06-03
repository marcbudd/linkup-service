package initalizers

import "github.com/marcbudd/linkup-service/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Comment{})
	DB.AutoMigrate(&models.Follow{})
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.Token{})
	DB.AutoMigrate(&models.User{})
}
