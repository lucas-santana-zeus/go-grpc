package routes

import (
	"go-grpc/client/handlers"
	"go-grpc/commons"
	"log"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()

	r.GET("/blocks/:id", handlers.GetBlockById)

	err := r.Run(*commons.PORTGinAPI)
	if err != nil {
		log.Fatalln("router error:", err)
	}
}
