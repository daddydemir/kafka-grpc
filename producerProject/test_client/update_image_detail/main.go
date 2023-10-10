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
	req := &proto.ImageRequest{
		ImageId:      "e7b2089a-5724-40fc-b3c2-a576ebd35247",
		Joy:          1,
		Anger:        2,
		Sorrow:       3,
		Surprise:     4,
		UnderExposed: 5,
		Blurred:      6,
		Confidence:   7,
		Headwear:     8,
	}

	resp, err := client.UpdateImageDetail(ctx, req)
	if err != nil {
		log.Fatalf("Failed to get image detail: %v", err)
	}

	fmt.Printf("Image: %v\n", resp)
}
