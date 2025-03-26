package controllers

import (
	"FinalCrossplatform/database"
	models "FinalCrossplatform/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// รับข้อมูล JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// ค้นหาลูกค้าจากฐานข้อมูลโดยใช้ email
	var customer models.Customer
	if err := database.DB.Where("email = ?", input.Email).First(&customer).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	// เปรียบเทียบรหัสผ่านตรงๆ (ไม่แฮช)
	if strings.TrimSpace(customer.Password) != strings.TrimSpace(input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(input.Password)); err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
	// 	return
	// }

	// หากเข้าสู่ระบบสำเร็จ
	c.JSON(http.StatusOK, gin.H{
		"message":     "Successfully logged in",
		"customer_id": customer.CustomerID,
		"first_name":  customer.FirstName,
		"last_name":   customer.LastName,
		"email":       customer.Email,
	})
}
