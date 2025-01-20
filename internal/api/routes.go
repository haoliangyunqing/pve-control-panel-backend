package api

import (
	"github.com/gin-gonic/gin"
	"pve-control-panel-backend/internal/models"
	"pve-control-panel-backend/internal/pve"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/pve/login", func(c *gin.Context) {
		var loginReq models.PVELoginRequest
		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(400, gin.H{"error": "invalid request body"})
			return
		}
		pveClient := pve.NewPVEClient(&pve.PVEConfig{
			Host:     loginReq.Host,
			Username: loginReq.Username,
			Password: loginReq.Password,
			Realm:    loginReq.Realm,
		})
		if err := pveClient.Authenticate(); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"message": "Login successful",
			"ticket":  pveClient.GetTicket(),
			"csrf":    pveClient.GetCSRFToken(),
		})
	})
}
