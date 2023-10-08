package service

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"image"
	"image/jpeg"
	"log"
	"makromusic/config"
	"makromusic/pkg/database"
	"makromusic/pkg/models"
	"os"
)

func GetAllImages() {
	var images []models.Image
	_ = database.DB.Find(&images)
	fmt.Println(images)
}

func SaveImage() {
	var image models.Image
	image.ImagePath = "C:\\Users\\demir\\Pictures\\chaos.png"
	image.ImageId = uuid.New()
	_ = database.DB.Save(&image)

}

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
	database.DB.Save(&img)
	AddQueue(img)

}
