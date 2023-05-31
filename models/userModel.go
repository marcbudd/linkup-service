package models

type User struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
}

type UserSignupRequestDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// func UserAutoMigrate() {
// 	db := initalizers.DB

// 	db.AutoMigrate(&User{})
// }

// func (u *User) SetPassword(password string) error {
// 	if len(password) == 0 {
// 		return errors.New("password cannot be empty")
// 	}
// 	bytePassword := []byte(password)
// 	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

// 	if err != nil {
// 		return errors.New("failed to hash password")
// 	}

// 	u.PasswordHash = string(passwordHash)
// 	return nil
// }

// func (u *User) CheckPassword(password string) error {
// 	bytePassword := []byte(password)
// 	byteHashedPassword := []byte(u.PasswordHash)
// 	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
// }

// func GetOneUser(condition interface{}) (User, error) {
// 	db := initalizers.DB
// 	var user User
// 	err := db.Where(condition).First(&user).Error
// 	return user, err
// }

// func CreateUser(data interface{}) error {
// 	db := initalizers.DB
// 	err := db.Save(data).Error
// 	return err
// }

// func (model *User) UpdateUser(data interface{}) error {
// 	db := initalizers.DB
// 	err := db.Model(model).Updates(data).Error
// 	return err
// }
