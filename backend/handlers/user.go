package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"backend/setting"

	"github.com/gin-gonic/gin"
)

type VPNUserInfo struct {
	Remark     string `json:"remark"`
	DataLeft   int64  `json:"data_left"`
	ExpireDate string `json:"expire_date"`
}

func GetUserInfo(c *gin.Context) {
	Remark := c.Param("remark")
	apiURL := fmt.Sprintf("%s/?remark=%s&apikey=%s",
		setting.AppConfig.VPNApiBaseURL,
		Remark,
		setting.AppConfig.VPNApiKey,
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Println("خطا در اتصال به API وی‌پی‌ان:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت اطلاعات"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "خطا از سمت سرور وی‌پی‌ان"})
		return
	}

	var userInfo VPNUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در پردازش داده‌ها"})
		return
	}

	c.JSON(http.StatusOK, userInfo)
}
