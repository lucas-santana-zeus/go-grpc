package main

import (
	"context"
	block "go-grpc/commons/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}

	client := block.NewBlocksClient(conn)
	req := block.RequestID{
		Id: "F1",
	}

	res, err := client.GetBlockById(context.Background(), &req)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res.GetId(), res.GetParentId())
}
