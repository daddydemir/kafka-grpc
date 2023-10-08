package models

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image"
	"image/jpeg"
	"io"
	"log"
	"makromusic/pkg/models"
	"os"
)

type RequestModel struct {
	Requests []Requests `json:"requests"`
}

type Requests struct {
	Image    `json:"image"`
	Features []Features `json:"features"`
}

type Image struct {
	Content string `json:"content"`
}
type Features struct {
	MaxResults int    `json:"maxResults"`
	Type       string `json:"type"`
}

func NewRequest(image *models.Image) io.Reader {
	request := RequestModel{
		Requests: []Requests{
			{
				Image: Image{
					Content: prepareRequest(image.ImagePath),
				},
				Features: []Features{
					{
						MaxResults: 50,
						Type:       "FACE_DETECTION",
					},
				},
			},
		},
	}
	arr, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return bytes.NewBuffer(arr)
}

func prepareRequest(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("file open error: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("image decode error: %v", err)
	}

	buf := new(bytes.Buffer)
	if err = jpeg.Encode(buf, img, nil); err != nil {
		log.Fatalf("jpeg encode error: %v", err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
