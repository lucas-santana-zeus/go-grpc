package handlers

import (
	"go-grpc/client/service"
	"go-grpc/commons/models"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func GetBlockByIdHandler(c *gin.Context) {
	id := c.Param("id")

	blockDTO, err := service.GetBlockById(id)

	if err != nil {
		log.Println("error block returning:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if reflect.DeepEqual(blockDTO, models.Block{}) {
		log.Println("empty block:", err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, blockDTO)
}
