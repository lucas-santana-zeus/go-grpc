package testUtils

import (
	"go-grpc/commons"
	"go-grpc/commons/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func SetupGrpcTestClient() block.BlocksClient {
	conn, err := grpc.Dial(*commons.PORTClient, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error dial create test connection", err.Error())
	}
	client := block.NewBlocksClient(conn)
	return client
}
