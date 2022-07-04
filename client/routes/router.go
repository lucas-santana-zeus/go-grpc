package routes

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRoutes() {
	r := gin.Default()

	r.GET("/blocks/:id")

	err := r.Run(":8000")
	if err != nil {
		log.Panicln(err)
	}
}
