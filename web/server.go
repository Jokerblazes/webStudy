package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Get user value
	r.GET("/car/:number", func(c *gin.Context) {
		number := c.Params.ByName("number")
		c.JSON(http.StatusOK, gin.H{"number": number, "name": "车名"})
	})

	return r
}
