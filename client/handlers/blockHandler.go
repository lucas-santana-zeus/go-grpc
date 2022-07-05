package handlers

import (
	"go-grpc/client/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBlockById(c *gin.Context) {
	id := c.Param("id")

	blocks, err := models.GetBlockById(id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"block:": blocks,
		})
	}

}
