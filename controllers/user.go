package controllers

import (
	"Template-Go/config"
	"Template-Go/models"
	"Template-Go/models/request"
	"Template-Go/models/response"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	req.Email = strings.TrimSpace(req.Email)
	req.Phone = strings.TrimSpace(req.Phone)

	var existingUser models.User
	if err := config.DB.Where("username = ? OR email = ?", req.Username, req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "User with this username or email already exists",
		})
		return
	}

	randomNumber := rand.Intn(9000) + 1000
	fiveMinutesLater := time.Now().Add(5 * time.Minute)
	user := models.User{
		Status:      "inactive",
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		Phone:       req.Phone,
		OtpVerify:   randomNumber,
		OtpReminder: &fiveMinutesLater,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, response.RegisterResponse{
		Message: "User registered successfully",
		User:    user,
	})
}
