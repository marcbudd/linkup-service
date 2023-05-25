package models

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	UserId  uint
	TweetId uint
}
