package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
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

	// C:\Users\demir\Downloads\mask.jpg
	imageData, err := ioutil.ReadFile("C:\\Users\\demir\\Pictures\\chaos.jpg") // Yüklenecek resmin dosya yolunu ayarlayın.
	if err != nil {
		log.Fatalf("Failed to read image file: %v", err)
	}

	ctx := context.Background()
	req := &proto.Request{
		ImageData: imageData,
	}

	resp, err := client.UploadImage(ctx, req)
	if err != nil {
		log.Fatalf("Failed to upload image: %v", err)
	}

	fmt.Printf("Uploaded image URL: %s\n", resp.Message)
}
