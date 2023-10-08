package api

import (
	"encoding/json"
	"io"
	"log"
	"makromusic/config"
	em "makromusic/external/models"
	"makromusic/pkg/models"
	"net/http"
)

func SendAnalyseRequest(image *models.Image) {

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
	var response = new(em.ResponseModel)
	err = json.Unmarshal(body, response)
	if err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)
	}
	log.Printf("Response: %v", response)
}

func Send() {}
