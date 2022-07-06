package handlers

import (
	"go-grpc/client/service"
	"go-grpc/commons/models"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func GetBlockById(c *gin.Context) {
	id := c.Param("id")

	blockDTO, err := service.GetBlockById(id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if reflect.DeepEqual(blockDTO, models.Block{}) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, blockDTO)
}
