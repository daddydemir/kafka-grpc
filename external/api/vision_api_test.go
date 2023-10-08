package api

import (
	"encoding/json"
	"log"
	"makromusic/external/models"
	"testing"
)

func TestSend(t *testing.T) {

	var f = new(models.Features)
	var i = new(models.Image)
	var r = new(models.Requests)
	var req []models.Requests
	var model models.RequestModel

	f.Type = "SAFE_SEARCH_DETECTION"
	f.MaxResults = 50
	i.Content = "base64"
	r.Image = *i
	r.Features = append(r.Features, *f)

	req = append(req, *r)
	model.Requests = req

	marshal, err := json.Marshal(model)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("json:%s ", marshal)
}
