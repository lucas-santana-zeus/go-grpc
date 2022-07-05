package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	block "go-grpc/commons/pb"
	"go-grpc/service/models"
	"google.golang.org/grpc"
	"net/http"
)

func GetBlockById(c *gin.Context) {
	id := c.Param("id")
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		//	todo: 1 - erro caso falhe a conexão com grpc client - {error: mgs} status 500
	}

	client := block.NewBlocksClient(conn)
	req := block.RequestID{
		Id: id,
	}

	res, err := client.GetBlockById(context.Background(), &req)
	if err != nil {
		//	todo: 2 - erro caso falhe a conexão com grpc client - {error: mgs} status 500
	}
	var blockDTO models.Block
	blockDTO.ID = res.GetId()
	blockDTO.ParentID = res.GetParentId()
	c.JSON(http.StatusOK, blockDTO)
}
