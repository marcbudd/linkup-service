package models

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FollwerUserId  uint //User Id of the user who is following
	FollowedUserId uint //User Id of the user who is being followed
}
