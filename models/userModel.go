package models

import "time"

type User struct {
	ID             uint      `gorm:"primary_key"`
	Username       string    `gorm:"column:username;not null;unique_index"`
	Email          string    `gorm:"column:email;not null;unique_index"`
	PasswordHash   string    `gorm:"column:password;not null"`
	EmailConfirmed bool      `gorm:"column:active;default:false;not null"`
	BirthDate      time.Time `gorm:"column:birth_date"`
	Name           string    `gorm:"column:name"`
	Bio            string    `gorm:"column:bio;size:1024"`
	Image          *string   `gorm:"column:image"`
}

type UserCreateRequestDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdatePasswortRequestDTO struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type UsersGetRequestDTO struct {
	Query string `form:"query"`
	Page  int    `form:"query"`
	Limit int    `form:"limit"`
}

type UserGetResponseDTO struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	BirthDate time.Time `json:"birthDate"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	Image     *string   `json:"image"`
}

type UserUpdateRequestDTO struct {
	Username  string    `json:"username"`
	BirthDate time.Time `json:"birthDate"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	Image     *string   `json:"image"`
}

type UserUpdatePasswordForgottenRequestDTO struct {
	Email string `json:"email"`
}
