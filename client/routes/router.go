package routes

import (
	"go-grpc/client/handlers"
	"go-grpc/commons"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()

	r.GET(*commons.ROUTEApi+":id", handlers.GetBlockByIdHandler)
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	err := r.Run(*commons.PORTGinAPI)
	if err != nil {
		log.Fatalln("router error:", err)
	}
}
