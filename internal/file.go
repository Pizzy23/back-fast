package handlers

import (
	models "fast/model"
	services "fast/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

type DataStandardizedRequest struct {
	PrfCadCodRnp string `json:"prfCadCodRnp"`
}

func GetDataStandardized(c *gin.Context) {
	prfCadCodRnp := c.Query("prfCadCodRnp")
	if prfCadCodRnp == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prfCadCodRnp is required"})
		return
	}

	creaResponse, err := services.SendCreaRequest(prfCadCodRnp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, creaResponse)
}

func PostDataStandardized(c *gin.Context) {
	var request models.CreaData
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creaResponse, err := services.SendCreaRequestPost(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	if _, err := xormEngine.Insert(creaResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, creaResponse)
}

func PutDataStandardized(c *gin.Context) {
	var request models.CreaData
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creaResponse, err := services.SendCreaRequestPost(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	if _, err := xormEngine.ID(request.PrfCrtCod).Update(creaResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, creaResponse)
}
