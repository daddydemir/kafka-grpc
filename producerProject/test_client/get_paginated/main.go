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

	resp, err := client.GetImageFeed(ctx, &proto.PaginationRequest{Limit: 2, Page: 2})
	if err != nil {
		log.Fatalf("Failed to get image feed: %v", err)
	}

	fmt.Printf("Response: %v\n", resp)
}
