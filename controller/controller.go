// Controller interface for the controller package.

package controller

import (
	"github.com/gin-gonic/gin"
)

//Controller ...
type Controller interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	GetById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
