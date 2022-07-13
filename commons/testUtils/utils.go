// Package testUtils contains the test utils for both grpc client and service
package testUtils

import (
	"go-grpc/commons"
	"go-grpc/commons/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// SetupGrpcTestClient setup a grpc client connection for unity test purpose
func SetupGrpcTestClient() block.BlocksClient {
	conn, err := grpc.Dial(*commons.PORTClient, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error dial create test connection", err.Error())
	}
	client := block.NewBlocksClient(conn)
	return client
}
