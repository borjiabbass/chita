package handlers

import (
	"backend/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionInput struct {
	Username      string `json:"username" binding:"required"`
	PlanID        int    `json:"plan_id" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

func CreateTransaction(c *gin.Context) {
	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر"})
		return
	}

	result, err := database.DB.Exec(
		"INSERT INTO transactions (username, plan_id, payment_method, status) VALUES (?, ?, ?, 'pending')",
		input.Username, input.PlanID, input.PaymentMethod,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ذخیره تراکنش"})
		return
	}

	id, _ := result.LastInsertId()

	c.JSON(http.StatusOK, gin.H{
		"status":         "pending",
		"message":        "تراکنش ذخیره شد",
		"transaction_id": id,
	})
}
