package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (controller *DefaultController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GET"})
}

func (controller *DefaultController) GetById(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

func (controller *DefaultController) Create(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

func (controller *DefaultController) Update(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

func (controller *DefaultController) Delete(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}
