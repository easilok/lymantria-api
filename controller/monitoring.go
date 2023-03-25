package controller

import (
	"net/http"

	"github.com/easilok/lymantria-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// GET /monitoring
// Get registered monitorings
func (h *BaseHandler) GetMonitoring(c *gin.Context) {
	_, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	var monitorings []models.Monitoring
	h.db.Model(&models.Monitoring{}).Where("deleted_at IS NULL").Preload("Butterflies.Register").Preload("Butterflies.Butterfly.Details").Preload(clause.Associations).Find(&monitorings)

	c.JSON(http.StatusOK, gin.H{"records": monitorings})
}

// GET /monitoring/:monitoringId/butterflies
// Get registered monitorings
func (h *BaseHandler) GetMonitoringButterflies(c *gin.Context) {
	_, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	monitoringId := c.Param("monitoringId")
	var monitoring models.Monitoring
	err := h.db.Model(&models.Monitoring{}).Where("id = ?", monitoringId).First(&monitoring).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"records": monitoring})
}
