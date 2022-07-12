package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"go-grpc/commons"
	block "go-grpc/commons/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	block.UnimplementedBlocksServer
	//	adicionar bqClient
	BQClient *bigquery.Client
}

//func (server *Server) mustEmbedUnimplementedBlocksServer() {}

func main() {
	server := grpc.NewServer()
	srv := &Server{}
	var err error
	srv.BQClient, err = bigquery.NewClient(context.Background(), "athena-dsv")
	if err != nil {
		log.Panicln(err)
	}
	block.RegisterBlocksServer(server, srv)

	listener, err := net.Listen("tcp", *commons.PORTService)
	if err != nil {
		log.Panicln(err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Panicln(err)
	}
}
