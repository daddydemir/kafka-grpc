package api

import (
	"consumerProject/config"
	em "consumerProject/external/models"
	"consumerProject/pkg/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Service interface {
	SaveAnalyse(analyse models.FacialAnalyse)
	UpdateImageStatus(id, text string)
}

func SendAnalyseRequest(image *models.Image, service Service) {

	request, err := http.NewRequest(http.MethodPost, config.GetEnv("VISION_API_URL"), em.NewRequest(image))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	var response *em.ResponseModel
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)
	}

	if !reflect.ValueOf(response.Response[0]).IsZero() {
		analyse := response.Response[0].FaceAnnotations[0].ToDBModel(image.ImageId)
		log.Printf("Response: %v", analyse)
		analyse.CreateDate = time.Now()
		service.SaveAnalyse(analyse)
		service.UpdateImageStatus(image.ImageId.String(), "ANALYSED")
		log.Printf("Analyse: %v", analyse)
	}
}
