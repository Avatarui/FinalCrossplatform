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

type PasswordChangeRequest struct {
	Email           string `json:"email"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

func ChangePassword(c *gin.Context) {
	var req PasswordChangeRequest

	// รับข้อมูล JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// ตรวจสอบความถูกต้องของข้อมูลนำเข้า
	if strings.TrimSpace(req.Email) == "" ||
		strings.TrimSpace(req.CurrentPassword) == "" ||
		strings.TrimSpace(req.NewPassword) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// ค้นหาลูกค้าในฐานข้อมูลโดยใช้อีเมล
	var customer models.Customer
	if err := database.DB.Where("email = ?", req.Email).First(&customer).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Customer not found"})
		return
	}

	// ตรวจสอบรหัสผ่านปัจจุบัน
	if strings.TrimSpace(customer.Password) != strings.TrimSpace(req.CurrentPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Current password is incorrect"})
		return
	}

	// ตรวจสอบความยาวรหัสผ่านใหม่
	if len(req.NewPassword) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password must be at least 8 characters long"})
		return
	}

	// อัปเดตรหัสผ่านในฐานข้อมูล
	customer.Password = req.NewPassword
	if err := database.DB.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update password"})
		return
	}

	// ส่งการตอบกลับสำเร็จ
	c.JSON(http.StatusOK, gin.H{
		"message": "Password changed successfully",
		"email":   customer.Email,
	})
}
