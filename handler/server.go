package handler

import (
	"context"
	"errors"
	"makromusic/pkg/service"
	"makromusic/proto"
)

type Server struct {
}

func (s *Server) UploadImage(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	if len(request.ImageData) >= 4*1024*1024 {
		return &proto.Response{
			Message: "Image is too large. maximum size 4MB",
		}, errors.New("image is too large. maximum size 4MB")
	}
	service.SaveAndWriteQueue(request.ImageData)
	return &proto.Response{Message: "Upload image is success."}, nil

}
