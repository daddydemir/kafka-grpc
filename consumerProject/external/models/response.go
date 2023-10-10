package models

import (
	em "consumerProject/pkg/models"
	"github.com/google/uuid"
)

type ResponseModel struct {
	Response []Response `json:"responses"`
}

type Response struct {
	FaceAnnotations []FaceAnnotations `json:"faceAnnotations"`
}

type FaceAnnotations struct {
	JoyLikelihood          string  `json:"joyLikelihood"`
	SorrowLikelihood       string  `json:"sorrowLikelihood"`
	AngerLikelihood        string  `json:"angerLikelihood"`
	SurpriseLikelihood     string  `json:"surpriseLikelihood"`
	UnderExposedLikelihood string  `json:"underExposedLikelihood"`
	BlurredLikelihood      string  `json:"blurredLikelihood"`
	HeadwearLikelihood     string  `json:"headwearLikelihood"`
	DetectionConfidence    float64 `json:"detectionConfidence"`
}

func (f *FaceAnnotations) ToDBModel(imageID uuid.UUID) em.FacialAnalyse {

	return em.FacialAnalyse{
		ImageID:      imageID,
		Joy:          getPoint(f.JoyLikelihood),
		Sorrow:       getPoint(f.SorrowLikelihood),
		Anger:        getPoint(f.AngerLikelihood),
		Surprise:     getPoint(f.SurpriseLikelihood),
		UnderExposed: getPoint(f.UnderExposedLikelihood),
		Blurred:      getPoint(f.BlurredLikelihood),
		Headwear:     getPoint(f.HeadwearLikelihood),
		Confidence:   f.DetectionConfidence * 100,
	}
}

func getPoint(val string) int {
	switch val {
	case "VERY_LIKELY":
		return 5
	case "LIKELY":
		return 4
	case "POSSIBLE":
		return 3
	case "UNLIKELY":
		return 2
	case "VERY_UNLIKELY":
		return 1
	}
	return 0
}
