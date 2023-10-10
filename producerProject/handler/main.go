package handler

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"producerProject/proto"
)

func Handler() {

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()

	proto.RegisterImageAnalyzeServiceServer(server, ImageAnalyseServer{})
	reflection.Register(server)
	fmt.Println("Server is running on port :5001...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
