package models

import (
	"context"
	block "go-grpc/commons/pb"
	"go-grpc/service/models"
	"log"

	"google.golang.org/grpc"
)

func GetBlockById(id string) (*models.Block, error) {

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
		return &models.Block{}, err
	}

	client := block.NewBlocksClient(conn)
	req := block.RequestID{
		Id: id,
	}

	res, err := client.GetBlockById(context.Background(), &req)

	if err != nil {
		log.Println(err.Error())
		return &models.Block{}, err
	}

	var blockDTO models.Block
	blockDTO.ID = res.GetId()
	blockDTO.ParentID = res.GetParentId()

	return &blockDTO, nil
}
