package services

import (
	"time"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/utils"
)

func CreateToken(userID uint) (*models.Token, error) {
	tokenString, err := utils.GenerateTokenString(20)
	if err != nil {
		return nil, err
	}

	db := initalizers.DB
	token := models.Token{
		ExpirationDate: time.Now().Add(time.Hour * 24 * 7),
		UserID:         userID,
		TokenString:    tokenString,
	}
	result := db.Create(&token)

	return &token, result.Error
}

func GetTokenByTokenString(tokenString string) (*models.Token, error) {
	db := initalizers.DB
	var token models.Token
	err := db.Where("token_string = ?", tokenString).First(&token).Error

	return &token, err
}

func ConfirmToken(token *models.Token) error {
	db := initalizers.DB
	token.ConfirmationDate = time.Now()
	err := db.Save(&token).Error
	return err
}
