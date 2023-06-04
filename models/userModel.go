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

// function to convert user to response dto
// can be called everywhere, changes can be made in one place
func ConvertRequestDTOToUser(req UserCreateRequestDTO) *User {
	return &User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.Password,
	}
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

type UserUpdateRequestDTO struct {
	Username  string    `json:"username"`
	BirthDate time.Time `json:"birthDate"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	Image     *string   `json:"image"`
}

// function to convert user to response dto
// can be called everywhere, changes can be made in one place
func (u *User) UpdateUser(req UserUpdateRequestDTO) {
	u.Username = req.Username
	u.BirthDate = req.BirthDate
	u.Name = req.Name
	u.Bio = req.Bio
	u.Image = req.Image
}

type UserUpdatePasswordForgottenRequestDTO struct {
	Email string `json:"email"`
}

type UserGetResponseDTO struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	BirthDate time.Time `json:"birthDate"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	Image     *string   `json:"image"`
}

// function to convert user to response dto
// can be called everywhere, changes can be made in one place
func (u *User) ConvertUserToResponseDTO() *UserGetResponseDTO {
	return &UserGetResponseDTO{
		ID:        u.ID,
		Username:  u.Username,
		BirthDate: u.BirthDate,
		Name:      u.Name,
		Bio:       u.Bio,
		Image:     u.Image,
	}
}
