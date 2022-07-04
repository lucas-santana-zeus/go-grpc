package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	block "go-grpc/commons/pb"
	"google.golang.org/grpc"
	"log"
)

func GetBlockById(c *gin.Context) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		//	todo: 1 - erro caso falhe a conexão com grpc client - {error: mgs} status 500
	}

	client := block.NewBlocksClient(conn)
	req := block.RequestID{
		Id: "F1",
	}

	res, err := client.GetBlockById(context.Background(), &req)
	if err != nil {
		//	todo: 2 - erro caso falhe a conexão com grpc client - {error: mgs} status 500
	}
	log.Println(res.GetId(), res.GetParentId())
}
