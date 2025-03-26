package controllers

import (
	"FinalCrossplatform/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var customer models.customer
	if err := database.DB.Where("LOWER(email) = ?", input.Email).First(&user).Error; err != nil {
		fmt.Println("User not found for email:", input.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email", "details": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid  password"})
		return
	}

	// ส่งข้อมูล User กลับ โดยไม่ส่ง Password
	c.JSON(http.StatusOK, gin.H{
		"customer_id": user.ID,
		"first_name":  user.FirstName,
		"last_name":   user.LastName,
		"email":       user.Email,
	})
}
