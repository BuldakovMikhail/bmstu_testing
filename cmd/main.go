package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var adminEmail string
var adminPassword string
var verificationCodes = make(map[string]string)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error loading .env file: %v", err)
	//}

	adminEmail = os.Getenv("RECIPIENT_EMAIL_ADDRESS")
	adminPassword = os.Getenv("RECIPIENT_PASSWORD")

	r := gin.Default()

	r.GET("/status", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	r.POST("/login", loginHandler)
	r.POST("/verify", verifyLoginHandler)
	r.POST("/reset_password", resetPasswordHandler)
	r.POST("/verify_reset_password", verifyResetHandler)

	if err := r.Run(":8084"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func loginHandler(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	//adminEmail = adminEmail[:len(adminEmail)-1]
	fmt.Println(request.Email)
	fmt.Println(adminEmail)
	fmt.Println(request.Email != adminEmail)

	fmt.Println(len(request.Password))
	fmt.Println(len(adminPassword))
	fmt.Println(request.Password != adminPassword)

	fmt.Println(adminPassword[3])

	adminPassword = adminPassword[:len(request.Password)]

	if request.Email != adminEmail || request.Password != adminPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	if os.Getenv("TEST") != "" {
		code = fmt.Sprintf("%d", 123456)
	}

	verificationCodes[request.Email] = code

	sendEmail(code)

	c.JSON(http.StatusOK, gin.H{"message": "Verification code sent to email"})
}

func verifyLoginHandler(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	expectedCode, exists := verificationCodes[request.Email]
	if !exists || expectedCode != request.Code {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid verification code"})
		return
	}

	delete(verificationCodes, request.Email)

	c.JSON(http.StatusOK, gin.H{"message": "Auth success!"})
}

func resetPasswordHandler(c *gin.Context) {
	var request struct {
		Email       string `json:"email"`
		OldPassword string `json:"oldPassword"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if request.Email != adminEmail || request.OldPassword != adminPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	if os.Getenv("TEST") != "" {
		code = fmt.Sprintf("%d", 123456)
	}

	verificationCodes[request.Email] = code

	sendEmail(code)

	c.JSON(http.StatusOK, gin.H{"message": "Password reset code sent to email"})
}

func verifyResetHandler(c *gin.Context) {
	var request struct {
		Email       string `json:"email"`
		Code        string `json:"code"`
		NewPassword string `json:"newPassword"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	expectedCode, exists := verificationCodes[request.Email]
	if !exists || expectedCode != request.Code {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid verification code"})
		return
	}

	adminPassword = request.NewPassword

	delete(verificationCodes, request.Email)

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully!"})
}

func sendEmail(code string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	senderEmail := os.Getenv("SENDER_EMAIL_ADDRESS")
	senderPassword := os.Getenv("SENDER_EMAIL_PASSWORD")
	smtpHost := os.Getenv("SMTP_SERVER")
	smtpHost = smtpHost[:len(smtpHost)-1]
	smtpPort := 587
	recipientEmail := os.Getenv("RECIPIENT_EMAIL_ADDRESS")
	recipientEmail = recipientEmail[:len(recipientEmail)-1]

	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", "Your Verification Code")
	m.SetBody("text/plain", fmt.Sprintf("Your verification code is: %s", code))

	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Error sending verification code: %v", err)
	}
}
