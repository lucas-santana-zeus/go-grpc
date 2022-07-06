package service

import (
	"context"
	block "go-grpc/commons/pb"
	"go-grpc/service/models"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"google.golang.org/grpc"
)

func GetBlockById(id string) (models.Block, error) {

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err.Error())
		return models.Block{}, err
	}

	client := block.NewBlocksClient(conn)
	req := block.RequestID{
		Id: id,
	}

	res, err := client.GetBlockById(context.Background(), &req)

	if err != nil {
		log.Println(err.Error())
		return models.Block{}, err
	}

	blockDTO := models.TransformResponseIntoBlockDTO(res)

	return blockDTO, nil
}
