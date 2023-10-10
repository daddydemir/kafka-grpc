package service

import (
	"bytes"
	"github.com/google/uuid"
	"image"
	"image/jpeg"
	"log"
	"os"
	"producerProject/config"
	"producerProject/pkg/database"
	"producerProject/pkg/models"
)

func SaveAndWriteQueue(imgs []byte) {
	var img models.Image
	img.ImageId = uuid.New()
	filename := uuid.New().String() + `.jpg`
	img.ImagePath = config.GetEnv("LOCAL_PATH") + filename

	file, err := os.Create(img.ImagePath)
	if err != nil {
		log.Fatalln(`file create has occurred error:`, err)
	}
	defer file.Close()
	imageFile, _, err := image.Decode(bytes.NewReader(imgs))
	if err != nil {
		log.Fatalln(`your file is not image.`, err)
	}
	err = jpeg.Encode(file, imageFile, nil)
	if err != nil {
		log.Fatalln(`image can not write.`, err)
	}
	// save local file
	img.Status = "WAITING"
	database.DB.Save(&img)
	AddQueue(img)

}
