package handler

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"producerProject/pkg/models"
	"producerProject/pkg/service"
	"producerProject/proto"
)

type ImageAnalyseServer struct{}

func (s ImageAnalyseServer) UploadImage(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	if len(request.ImageData) >= 4*1024*1024 {
		return &proto.Response{
			Message: "Image is too large. maximum size 4MB",
		}, errors.New("image is too large. maximum size 4MB")
	}
	service.SaveAndWriteQueue(request.ImageData)
	return &proto.Response{Message: "Upload image is success."}, nil

}

func (s ImageAnalyseServer) GetImageDetail(ctx context.Context, request *proto.ImageDetailRequest) (*proto.ImageResponse, error) {
	img := service.GetImageByID(request.ImageId)

	var response = new(proto.ImageResponse)
	response.ImageId = img.ImageID.String()
	response.Joy = int32(img.Joy)
	response.Sorrow = int32(img.Sorrow)
	response.Anger = int32(img.Anger)
	response.Surprise = int32(img.Surprise)
	response.Headwear = int32(img.Headwear)
	response.UnderExposed = int32(img.UnderExposed)
	response.Blurred = int32(img.Blurred)
	response.Confidence = float32(img.Confidence)
	response.CreateDate = timestamppb.New(img.CreateDate)
	return response, nil
}

func (s ImageAnalyseServer) UpdateImageDetail(ctx context.Context, request *proto.ImageRequest) (*proto.Response, error) {
	var err error
	image := new(models.FacialAnalyse)
	image.ImageID, err = uuid.Parse(request.ImageId)
	if err != nil {
		log.Fatalf("Parse image id error: %v", err)
		return nil, err
	}
	image.Joy = int(request.Joy)
	image.Anger = int(request.Anger)
	image.Sorrow = int(request.Sorrow)
	image.Surprise = int(request.Surprise)
	image.UnderExposed = int(request.UnderExposed)
	image.Blurred = int(request.Blurred)
	image.Headwear = int(request.Headwear)
	image.Confidence = float64(request.Confidence)
	service.UpdateAnalyse(*image)
	return &proto.Response{Message: "Update image is success."}, nil
}

func (s ImageAnalyseServer) GetImageFeed(ctx context.Context, request *proto.PaginationRequest) (*proto.Images, error) {
	var result []models.FacialAnalyse
	var response []*proto.ImageResponse
	result = service.GetPaginatedResult(int(request.Limit), int(request.Page))
	for r := range result {
		var img proto.ImageResponse
		img.ImageId = result[r].ImageID.String()
		img.Joy = int32(result[r].Joy)
		img.Sorrow = int32(result[r].Sorrow)
		img.Anger = int32(result[r].Anger)
		img.Surprise = int32(result[r].Surprise)
		img.UnderExposed = int32(result[r].UnderExposed)
		img.Blurred = int32(result[r].Blurred)
		img.Headwear = int32(result[r].Headwear)
		img.Confidence = float32(result[r].Confidence)
		img.CreateDate = timestamppb.New(result[r].CreateDate)
		response = append(response, &img)
	}
	return &proto.Images{Images: response}, nil
}
