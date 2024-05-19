package main

import (
	"example/SES4Case/handlers"
	"example/SES4Case/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SendEmails(db *gorm.DB) {
	var emails []models.Email
	result := db.Find(&emails)
	if result.Error != nil {
		log.Printf("Error fetching emails: %v", result.Error)
		return
	}

	rate, err := handlers.FetchRate()
	if err != nil {
		log.Printf("Error fetching rate: %v", err)
		return
	}

	for _, email := range emails {
		SendEmail(email.Address, rate)
	}
}

func SendEmail(to string, rate float64) {
	m := gomail.NewMessage()
	m.SetHeader("From", "evh.vorobiov@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Daily USD to UAH Rate")
	m.SetBody("text/plain", fmt.Sprintf("Today's USD to UAH rate is: %.2f", rate))

	app_pass := os.Getenv("GM_APP_PASSWORD")
	d := gomail.NewDialer("smtp.gmail.com", 587, "evh.vorobiov@gmail.com", app_pass)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error sending email to %s: %v", to, err)
	}
}

func main() {
	connStr := os.Getenv("DB_CONN")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&models.Email{})

	router := gin.Default()

	router.GET("/api/rate", handlers.GetRate)
	router.POST("/api/subscribe", handlers.Subscribe(db))

	go func() {
		for {
			SendEmails(db)
			time.Sleep(24 * time.Hour)
		}
	}()

	router.Run(":8080")
}
