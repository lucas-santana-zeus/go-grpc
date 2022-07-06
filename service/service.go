package main

import (
	"context"
	"go-grpc/commons"
	"go-grpc/commons/models"
	block "go-grpc/commons/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	block.UnimplementedBlocksServer
}

func (server *Server) GetBlockById(context context.Context, req *block.RequestID) (*block.ResponseBlock, error) {
	blockDAO := models.GetBlockById(req.GetId())
	resBlock := models.TransformBlockDAOIntoResponse(blockDAO)
	return resBlock, nil
}

//func (server *Server) mustEmbedUnimplementedBlocksServer() {}

func main() {
	server := grpc.NewServer()
	block.RegisterBlocksServer(server, &Server{})

	listener, err := net.Listen("tcp", *commons.PORTService)
	if err != nil {
		log.Panicln(err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Panicln(err)
	}
}
