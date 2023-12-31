package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"producerProject/proto"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewImageAnalyzeServiceClient(conn)

	ctx := context.Background()
	req := &proto.ImageDetailRequest{
		ImageId: "0f536854-0ffa-46b9-aa1a-660c0b16d0ad",
	}

	resp, err := client.GetImageDetail(ctx, req)
	if err != nil {
		log.Fatalf("Failed to get image detail: %v", err)
	}

	fmt.Printf("Image: %v\n", resp)
}
