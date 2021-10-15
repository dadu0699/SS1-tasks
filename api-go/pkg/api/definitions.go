package api

// Response is the structure of the request response.
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// DetectLabelsStruct is the structure of the body sent to the request.
type DetectLabelsStruct struct {
	Image string         `json:"image"`
}

// CompareFacesStruct is the structure of the body sent to the request.
type CompareFacesStruct struct {
	FirstImage string         `json:"imagen1"`
	SecondImage string         `json:"imagen2"`
	Similarity float64         `json:"similitud"`
}
