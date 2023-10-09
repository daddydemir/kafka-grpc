package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"makromusic/proto"
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
		ImageId: "e7b2089a-5724-40fc-b3c2-a576ebd35247",
	}

	resp, err := client.GetImageDetail(ctx, req)
	if err != nil {
		log.Fatalf("Failed to get image detail: %v", err)
	}

	fmt.Printf("Image: %v\n", resp)
}