package handler

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"makromusic/proto"
	"net"
)

func Handler() {

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterImageAnalyzeServiceServer(server, &Server{})
	fmt.Println("Server is running on port :5001...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
