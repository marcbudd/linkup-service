package initalizers

import "github.com/marcbudd/twitter-clone-backend/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
