package handlers

import (
	models "fast/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func GetShareData(c *gin.Context) {
	engine, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	xormEngine, ok := engine.(*xorm.Engine)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection invalid"})
		return
	}

	var shareData []models.ShareData
	if err := xormEngine.Find(&shareData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shareData)
}

func PostShareData(c *gin.Context) {
	var request models.ShareData
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	engine, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	xormEngine, ok := engine.(*xorm.Engine)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection invalid"})
		return
	}

	if _, err := xormEngine.Insert(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, request)
}

func PutShareData(c *gin.Context) {
	var request models.ShareData
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	engine, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	xormEngine, ok := engine.(*xorm.Engine)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection invalid"})
		return
	}

	if _, err := xormEngine.ID(request.ID).Update(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}
