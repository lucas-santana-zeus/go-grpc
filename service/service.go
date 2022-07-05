package main

import (
	"context"
	block "go-grpc/commons/pb"
	"go-grpc/service/models"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	block.UnimplementedBlocksServer
}

func (server *Server) GetBlockById(context context.Context, req *block.RequestID) (*block.ResponseBlock, error) {
	resBlock := models.GetBlockById(req.GetId())
	return resBlock, nil
}
func (server *Server) mustEmbedUnimplementedBlocksServer() {}

func main() {
	server := grpc.NewServer()
	block.RegisterBlocksServer(server, &Server{})

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Panicln(err)
	}
}
