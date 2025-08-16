package main

import (
	"backend/database"
	"backend/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// اتصال به دیتابیس
	if err := database.Connect(); err != nil {
		log.Fatal("خطا در اتصال به دیتابیس:", err)
	}

	r := gin.Default()

	// مسیر تعرفه‌ها
	r.GET("/plans", handlers.GetPlans)

	r.Run(":8080")
}
