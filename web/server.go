package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webStudy/car"
)

type Server struct {
	Engine *gin.Engine
	Car    car.Service
}

func NewServer(car car.Service) *Server {
	server := &Server{Car: car}
	server.Engine = server.setupRouter()
	return server
}

func (server *Server) setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/cars", func(c *gin.Context) {
		number := c.Params.ByName("number")
		c.JSON(http.StatusOK, gin.H{"number": number, "name": "车名"})
	})

	return r
}
