package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
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
	fmt.Println("Created bqclient")
	if err != nil {
		log.Fatalln("bqclient instance error:", err)
	}
	block.RegisterBlocksServer(server, srv)
	fmt.Println("server registered")

	listener, err := net.Listen("tcp", *commons.PORTService)
	if err != nil {
		log.Fatalln("listener instance error:", err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Fatalln("server.serve error:", err)
	}
	fmt.Println("server running")
}
