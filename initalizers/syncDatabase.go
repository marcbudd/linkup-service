package initalizers

import "github.com/marcbudd/linkup-service/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.Like{})
	DB.AutoMigrate(&models.Follow{})
	DB.AutoMigrate(&models.Comment{})
}
