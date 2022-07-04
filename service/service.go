package main

import (
	"context"
	block "go-grpc/commons/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	block.UnimplementedBlocksServer
}

func (server *Server) GetBlockById(context context.Context, req *block.RequestID) (*block.ResponseBlock, error) {
	resBlock := &block.ResponseBlock{
		Id:       "F1",
		ParentId: "C0",
	}
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
