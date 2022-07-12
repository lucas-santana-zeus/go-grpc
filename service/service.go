package main

import (
	"cloud.google.com/go/bigquery"
	"go-grpc/commons"
	block "go-grpc/commons/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	block.UnimplementedBlocksServer
	BQClient *bigquery.Client
	//	adicionar bqClient
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
