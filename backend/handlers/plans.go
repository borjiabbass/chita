package handlers

import (
	"backend/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Plan struct {
	ID           int     `json:"id"`
	PlanName     string  `json:"plan_name"`
	DurationDays int     `json:"duration_days"`
	Price        float64 `json:"price"`
}

func GetPlans(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, plan_name, duration_days, price FROM vpn_plans")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در خواندن تعرفه‌ها"})
		return
	}
	defer rows.Close()

	var plans []Plan
	for rows.Next() {
		var p Plan
		if err := rows.Scan(&p.ID, &p.PlanName, &p.DurationDays, &p.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در پردازش داده‌ها"})
			return
		}
		plans = append(plans, p)
	}

	c.JSON(http.StatusOK, plans)
}
