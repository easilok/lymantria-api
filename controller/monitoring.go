package controller

import (
	"net/http"

	"github.com/easilok/lymantria-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// GET /monitoring
// Get registered monitorings
func (h *BaseHandler) GetAllMonitorings(c *gin.Context) {
	_, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	var monitorings []models.Monitoring
	h.db.Model(&models.Monitoring{}).Where("deleted_at IS NULL").Preload("Appearances.Register").Preload("Appearances.Butterfly.Details").Preload(clause.Associations).Find(&monitorings)

	c.JSON(http.StatusOK, gin.H{"records": monitorings})
}

// GET /monitoring/:monitoringId
// Get single monitoring and butterflies appearances
func (h *BaseHandler) GetMonitoring(c *gin.Context) {
	_, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	monitoringId := c.Param("monitoringId")
	var monitoring models.Monitoring
	err := h.db.Model(&models.Monitoring{}).Where("deleted_at IS NULL").Where("id = ?", monitoringId).Preload("Appearances.Register").Preload("Appearances.Butterfly.Details").Preload(clause.Associations).First(&monitoring).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"records": monitoring})
}

// GET /monitoring/latest
// Get latest monitoring and butterflies appearances
func (h *BaseHandler) GetMonitoringLatest(c *gin.Context) {
	var monitoring models.Monitoring
	err := h.db.Model(&models.Monitoring{}).Where("deleted_at IS NULL").Order("registered_at DESC").Preload("Appearances.Register").Preload("Appearances.Butterfly.Details").Preload(clause.Associations).First(&monitoring).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"records": monitoring})
}
