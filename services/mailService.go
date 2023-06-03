package services

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendTokenMail(receiver string, token string) {

	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	address := host + ":" + port
	link := "http://localhost:3000/confirmEmail/" + token

	to := []string{receiver}

	subject := "Subject: Welcome to LinkUp Social Media!\nPlease verify your email address"
	body := "Please verify your email address using the following link: \n" + link
	message := []byte(subject + "\n\n" + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func SendConfirmedEmail(receiver string) {
	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	address := host + ":" + port

	to := []string{receiver}

	subject := "Subject: Email confirmed\n"
	body := "Thank you! Your email address has been confirmed"
	message := []byte(subject + "\n\n" + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func SendPasswordForgottenMail(receiver string, newPassword string) {
	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	address := host + ":" + port

	to := []string{receiver}

	subject := "Subject: *Confidential* Password reset\n"
	body := "Your new password is: " + newPassword + "\nPlease change your password after logging in!"
	message := []byte(subject + "\n\n" + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)

	if err != nil {
		fmt.Println(err)
		return
	}
}
