package routes

import (
	"github.com/gin-gonic/gin"
	"go-grpc/client/handlers"
	"log"
)

func HandleRoutes() {
	r := gin.Default()

	r.GET("/blocks/:id", handlers.GetBlockById)

	err := r.Run(":8000")
	if err != nil {
		log.Panicln(err)
	}
}
