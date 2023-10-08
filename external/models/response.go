package models

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
