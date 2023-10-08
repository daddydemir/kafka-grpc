package handler

import (
	"bufio"
	"context"
	"errors"
	"log"
	"makromusic/config"
	"makromusic/pkg/service"
	"makromusic/proto"
	"os"
)

type Server struct {
}

func (s *Server) UploadImage(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	if len(request.ImageData) >= 4*1024*1024 {
		return &proto.Response{
			Message: "Image is too large. maximum size 4MB",
		}, errors.New("image is too large. maximum size 4MB")
	}
	//buffer := base64.StdEncoding.EncodeToString(request.ImageData)
	//test(buffer)
	service.SaveAndWriteQueue(request.ImageData)

	return &proto.Response{Message: "Upload image is success."}, nil

}

func test(s string) {

	file, _ := os.Create(config.GetEnv("LOCAL_PATH") + `file.txt`)

	writer := bufio.NewWriter(file)
	_, err := writer.WriteString(s)
	if err != nil {
		log.Fatalln(`write error:`, err)
	}
	writer.Flush()
}
