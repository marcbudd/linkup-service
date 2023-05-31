package models

import (
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	Following    User // User that is being followed
	FollowingID  uint
	FollowedBy   User // User that is following
	FollowedByID uint
}

type FollowCreateRequestDTO struct {
	UserID uint `json:"userId"`
}

// func FollowAutoMigrate() {
// 	db := initalizers.DB

// 	db.AutoMigrate(&Follow{})
// }

// // u wants to follow v
// func (u User) CreateFollow(v User) error {
// 	db := initalizers.DB
// 	var follow Follow
// 	err := db.FirstOrCreate(&follow, &Follow{
// 		FollowingID:  v.ID,
// 		FollowedByID: u.ID,
// 	}).Error
// 	return err
// }

// // is u following v?
// func (u User) IsFollowing(v User) bool {
// 	db := initalizers.DB
// 	var follow Follow
// 	db.Where(Follow{
// 		FollowingID:  v.ID,
// 		FollowedByID: u.ID,
// 	}).First(&follow)
// 	return follow.ID != 0
// }

// func (u User) DeleteFollow(v User) error {
// 	db := initalizers.DB
// 	err := db.Where(Follow{
// 		FollowingID:  v.ID,
// 		FollowedByID: u.ID,
// 	}).Delete(Follow{}).Error
// 	return err
// }

// func (u User) GetFollowings() []User {
// 	db := initalizers.DB

// 	var follows []Follow
// 	db.Find(&follows).Where("followed_by_id = ?", u.ID)

// 	var followings []User
// 	for _, follow := range follows {
// 		var temp User
// 		db.Find(&temp).Where("id = ?", follow.FollowingID)
// 		followings = append(followings, temp)
// 	}

// 	return followings
// }
