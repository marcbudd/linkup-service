package controllers

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func SendMail(c *gin.Context) {

	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")

	receiver := "marcbuddemeier@gmail.com"
	to := []string{receiver}

	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	address := host + ":" + port

	subject := "Go Test Mail"
	body := "This is a test mail"
	message := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)

	if err != nil {
		fmt.Println(err)
		return
	}

}
