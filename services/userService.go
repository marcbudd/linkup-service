package services

import (
	"errors"
	"time"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/utils"
)

func CreateUser(req models.UserCreateRequestDTO) (*models.User, error) {
	// Validate input
	if !utils.IsValidEmail(req.Email) {
		return nil, errors.New("email is not valid")
	}
	if !utils.IsValidUsername(req.Username) {
		return nil, errors.New("username is not valid")
	}
	if !utils.IsValidPassword(req.Password) {
		return nil, errors.New("password is not valid")
	}

	// Check if user exists
	db := initalizers.DB
	var count int64 = 0
	db.Model(&models.User{}).Where("email = ?", req.Email).Count(&count)
	if count > 0 {
		return nil, errors.New("email is already taken")
	}
	db.Model(&models.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		return nil, errors.New("username is already taken")
	}

	//Hash password
	passwordHashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	req.Password = passwordHashed

	// Create user
	user := *models.ConvertRequestDTOToUser(req)

	err = db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	// Create token and send mail
	token, err := CreateToken(user.ID)
	if err != nil {
		return nil, err
	}
	SendTokenMail(user.Email, token.TokenString)

	return &user, err
}

func LoginUser(req models.UserLoginRequestDTO) (*models.User, error) {
	// Find user by email
	db := initalizers.DB
	var user models.User
	err := db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return nil, errors.New("email or password is incorrect")
	}

	// Check password
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("email or password is incorrect")
	}

	return &user, nil
}

func ConfirmEmail(userID uint, tokenString string) error {
	db := initalizers.DB

	var user models.User
	db.Where("id = ?", userID).First(&user)

	token, err := GetTokenByTokenString(tokenString)

	// if token is expired or not found, send new token
	if err != nil || token.UserID != user.ID || token.ExpirationDate.After(time.Now()) {
		token, _ := CreateToken(userID)
		SendTokenMail(user.Email, token.TokenString)
		return errors.New("token is expired")
	}

	// Set email to confirmed
	ConfirmToken(token)
	user.EmailConfirmed = true
	db.Save(&user)
	SendConfirmedEmail(user.Email)

	return nil

}

func UpdatePassword(userID uint, req models.UserUpdatePasswortRequestDTO) error {

	// Find user by id
	db := initalizers.DB
	var user models.User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return errors.New("user not found")
	}

	// Check password
	if !utils.CheckPassword(req.OldPassword, user.PasswordHash) {
		return errors.New("old password is incorrect")
	}

	isStrong := utils.IsValidPassword(req.NewPassword)
	if !isStrong {
		return errors.New("password is not strong enough")
	}

	// Hash password
	passwordHashed, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Update password
	user.PasswordHash = passwordHashed
	db.Save(&user)

	return nil
}

func GetUserByID(id string) (*models.UserGetResponseDTO, error) {
	db := initalizers.DB
	var user models.User
	err := db.Where("id = ?", id).First(&user).Error

	var responseUser = *user.ConvertUserToResponseDTO()
	return &responseUser, err
}

func GetUsers(query string, page int, limit int) (*[]models.UserGetResponseDTO, error) {
	db := initalizers.DB
	var users []models.UserGetResponseDTO

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
	return &users, err
}

func UpdateUser(userID uint, req models.UserUpdateRequestDTO) error {

	// Check if username already exists
	var count int64 = 0
	db := initalizers.DB
	db.Model(&models.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		return errors.New("username is already taken")
	}

	// Find user by id
	var user models.User
	db.Where("id = ?", userID).First(&user)

	// Update user
	user.UpdateUser(req)

	err := db.Save(&user).Error

	return err

}

func UpdatePasswordForgotten(req models.UserUpdatePasswordForgottenRequestDTO) error {
	db := initalizers.DB
	var user models.User
	result := db.Where("email = ?", req.Email).First(&user)

	if result.Error != nil {
		return errors.New("user not found")
	}

	randomPassword, _ := utils.GenerateTokenString(20)
	passwordHashed, err := utils.HashPassword(randomPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.PasswordHash = passwordHashed
	db.Save(&user)
	SendPasswordForgottenMail(user.Email, randomPassword)

	return nil

}
