package services

import (
	"net/http"
	"time"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/linkuperrors"
	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/utils"
)

func CreateUser(req models.UserCreateRequestDTO) (*models.UserGetResponseDTO, *linkuperrors.LinkupError) {
	// Validate input
	if !utils.IsValidEmail(req.Email) {
		return nil, linkuperrors.New("email is not valid", http.StatusBadRequest)
	}
	if !utils.IsValidUsername(req.Username) {
		return nil, linkuperrors.New("username is not valid", http.StatusBadRequest)
	}
	if !utils.IsValidPassword(req.Password) {
		return nil, linkuperrors.New("password is not strong enough", http.StatusBadRequest)
	}

	// Check if user exists
	db := initalizers.DB
	var count int64 = 0
	db.Model(&models.User{}).Where("email = ?", req.Email).Count(&count)
	if count > 0 {
		return nil, linkuperrors.New("email is already taken", http.StatusConflict)
	}
	db.Model(&models.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		return nil, linkuperrors.New("username is already taken", http.StatusConflict)
	}

	//Hash password
	passwordHashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, linkuperrors.New("failed to hash password", http.StatusInternalServerError)
	}
	req.Password = passwordHashed

	// Create user
	user := *models.ConvertRequestDTOToUser(req)

	err = db.Save(&user).Error
	if err != nil {
		return nil, linkuperrors.New(err.Error(), http.StatusInternalServerError)
	}

	// Create token and send mail
	token, err := CreateToken(user.ID)
	if err != nil {
		return nil, linkuperrors.New(err.Error(), http.StatusInternalServerError)
	}

	err = SendTokenMail(user.Email, token.TokenString)
	if err != nil {
		return nil, linkuperrors.New(err.Error(), http.StatusInternalServerError)
	}

	return user.ConvertUserToResponseDTO(), nil
	// return nil, nil
}

func LoginUser(req models.UserLoginRequestDTO) (*models.User, *linkuperrors.LinkupError) {
	// Find user by email
	db := initalizers.DB
	var user models.User
	err := db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return nil, linkuperrors.New("email or password is incorrect", http.StatusUnauthorized)
	}

	// Check password
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, linkuperrors.New("email or password is incorrect", http.StatusUnauthorized)
	}

	return &user, nil
}

func ConfirmEmail(tokenString string) *linkuperrors.LinkupError {
	// Get token
	db := initalizers.DB
	token, err := GetTokenByTokenString(tokenString)

	// if token not found, send new token
	if err != nil || token == nil {
		return linkuperrors.New("token is expired or not found", http.StatusUnauthorized)
	}

	// Get user
	var user models.User
	err = db.Where("id = ?", token.UserID).First(&user).Error
	if err != nil {
		return linkuperrors.New("user not found", http.StatusNotFound)
	}

	// Check if token is expired
	if err != nil || token.ExpirationDate.After(time.Now()) {
		token, _ := CreateToken(token.UserID)
		SendTokenMail(user.Email, token.TokenString)
		return linkuperrors.New("token is expired or not found", http.StatusUnauthorized)
	}

	// Set token to confirmed and send email
	ConfirmToken(token)
	user.EmailConfirmed = true
	db.Save(&user)
	err = SendConfirmedEmail(user.Email)
	if err != nil {
		return linkuperrors.New(err.Error(), http.StatusInternalServerError)
	}

	return nil

}

func UpdatePassword(userID uint, req models.UserUpdatePasswortRequestDTO) *linkuperrors.LinkupError {

	// Find user by id
	db := initalizers.DB
	var user models.User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return linkuperrors.New("user not found", http.StatusNotFound)
	}

	// Check password
	if !utils.CheckPassword(req.OldPassword, user.PasswordHash) {
		return linkuperrors.New("old password is incorrect", http.StatusUnauthorized)
	}

	isStrong := utils.IsValidPassword(req.NewPassword)
	if !isStrong {
		return linkuperrors.New("password is not strong enough", http.StatusBadRequest)
	}

	// Hash password
	passwordHashed, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return linkuperrors.New("failed to hash password", http.StatusInternalServerError)
	}

	// Update password
	user.PasswordHash = passwordHashed
	err = db.Save(&user).Error

	if err != nil {
		return linkuperrors.New(err.Error(), http.StatusInternalServerError)
	}

	return nil
}

func GetUserByID(id string, currentUserID uint) (*models.UserDetailGetResponseDTO, *linkuperrors.LinkupError) {
	db := initalizers.DB
	var user models.User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, linkuperrors.New("user not found", http.StatusNotFound)
	}

	// Get number of followers
	var numberFollowers int64
	db.Model(&models.Follow{}).Where("user_followed_id = ?", id).Count(&numberFollowers)

	// Get number of following
	var numberFollowing int64
	db.Model(&models.Follow{}).Where("user_following_id = ?", id).Count(&numberFollowing)

	// Is user followed by logged in user
	isFollowing := false
	var count int64
	db.Model(&models.Follow{}).Where("user_followed_id = ? AND user_following_id = ?", id, currentUserID).Count(&count)
	if count > 0 {
		isFollowing = true
	}

	var responseUser = user.ConvertUserToDetailResponseDTO(numberFollowers, numberFollowing, isFollowing)
	return responseUser, nil
}

func GetUsers(query string, page int, limit int, currentUserID uint) (*[]models.UserDetailGetResponseDTO, *linkuperrors.LinkupError) {
	db := initalizers.DB
	var users []models.User

	// Set default values
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	// Database query with optional search parameter
	query = "%" + query + "%"
	dbQuery := db.Model(&models.User{}).
		Select("id, username, birth_date, name, bio, image").
		Where("username LIKE ? OR name LIKE ?", query, query)

	// Perform database query with pagination
	offset := (page - 1) * limit
	err := dbQuery.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, linkuperrors.New(err.Error(), http.StatusInternalServerError)
	}

	// Convert to response dtos
	var dtos []models.UserDetailGetResponseDTO
	for _, user := range users {
		// Get number of followers
		var numberFollowers int64
		db.Model(&models.Follow{}).Where("user_followed_id = ?", user.ID).Count(&numberFollowers)

		// Get number of following
		var numberFollowing int64
		db.Model(&models.Follow{}).Where("user_following_id = ?", user.ID).Count(&numberFollowing)

		// Is user followed by logged in user
		isFollowing := false
		var count int64
		db.Model(&models.Follow{}).Where("user_followed_id = ? AND user_following_id = ?", user.ID, currentUserID).Count(&count)
		if count > 0 {
			isFollowing = true
		}

		dtos = append(dtos, *user.ConvertUserToDetailResponseDTO(numberFollowers, numberFollowing, isFollowing))
	}

	return &dtos, nil
}

func UpdateUser(userID uint, req models.UserUpdateRequestDTO) *linkuperrors.LinkupError {

	// Check username
	if !utils.IsValidUsername(req.Username) {
		return linkuperrors.New("username is not valid", http.StatusBadRequest)
	}

	// Check if username already exists
	var count int64 = 0
	db := initalizers.DB
	db.Model(&models.User{}).Where("username = ? and user_id <> ?", req.Username, userID).Count(&count)
	if count > 0 {
		return linkuperrors.New("username is already taken", http.StatusConflict)
	}

	// Find user by id
	var user models.User
	db.Where("id = ?", userID).First(&user)

	// Update user
	user.UpdateUser(req)

	err := db.Save(&user).Error
	if err != nil {
		return linkuperrors.New(err.Error(), http.StatusInternalServerError)
	}

	return nil

}

func UpdatePasswordForgotten(req models.UserUpdatePasswordForgottenRequestDTO) *linkuperrors.LinkupError {
	db := initalizers.DB
	var user models.User
	result := db.Where("email = ?", req.Email).First(&user)

	if result.Error != nil {
		return linkuperrors.New("user not found", http.StatusNotFound)
	}

	randomPassword, _ := utils.GenerateTokenString(20)
	passwordHashed, err := utils.HashPassword(randomPassword)
	if err != nil {
		return linkuperrors.New("failed to hash password", http.StatusInternalServerError)
	}

	// Update Password
	user.PasswordHash = passwordHashed
	db.Save(&user)

	// Send new password
	err = SendPasswordForgottenMail(user.Email, randomPassword)
	if err != nil {
		return linkuperrors.New(err.Error(), http.StatusInternalServerError)
	}

	return nil

}

func DeleteUser(userID uint) *linkuperrors.LinkupError {
	db := initalizers.DB
	var user models.User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return linkuperrors.New("user not found", http.StatusNotFound)
	}

	// Delete all comments from user
	db.Where("user_id = ?", userID).Delete(&models.Comment{})

	// Delete all chats from user
	db.Where("sender_id = ? OR receiver_id = ?", userID, userID).Delete(&models.Message{})

	// Delete all followers from user
	db.Where("followed_id = ? OR follower_id = ?", userID, userID).Delete(&models.Follow{})

	// Delete all tokens from user
	db.Where("user_id = ?", userID).Delete(&models.Token{})

	// Delete all likes from user
	db.Where("user_id = ?", userID).Delete(&models.Like{})

	// Get all posts by user
	var posts []models.Post
	db.Where("user_id = ?", userID).Find(&posts)

	// Delete all comments and likes from posts
	for _, post := range posts {
		db.Where("post_id = ?", post.ID).Delete(&models.Comment{})
		db.Where("post_id = ?", post.ID).Delete(&models.Like{})
	}

	// Delete all posts
	db.Where("user_id = ?", userID).Delete(&models.Post{})

	// Delete user
	db.Delete(&user)

	return nil
}
