package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type myEvent struct {
	FirstImage  string  `json:"imagen1"`
	SecondImage string  `json:"imagen2"`
	Similarity  float64 `json:"similitud"`
}

func createSession() (*rekognition.Rekognition, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-2"),
		},
	)
	if err != nil {
		return nil, err
	}

	svc := rekognition.New(sess)
	return svc, nil
}

func compareFaces(body myEvent) (*rekognition.CompareFacesOutput, error) {
	decodedFirstImage, err := base64.StdEncoding.DecodeString(body.FirstImage)
	if err != nil {
		return nil, err
	}

	decodedSecondImage, err := base64.StdEncoding.DecodeString(body.SecondImage)
	if err != nil {
		return nil, err
	}

	input := &rekognition.CompareFacesInput{
		SimilarityThreshold: aws.Float64(body.Similarity),
		SourceImage:         &rekognition.Image{Bytes: decodedFirstImage},
		TargetImage:         &rekognition.Image{Bytes: decodedSecondImage},
	}

	svc, err := createSession()
	if err != nil {
		return nil, err
	}

	result, err := svc.CompareFaces(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func handleLambdaEvent(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "POST",
		},
	}

	req := myEvent{}
	err := json.Unmarshal([]byte(r.Body), &req)
	if err != nil {
		resp.Body = "Cannot parse post body"
		resp.StatusCode = http.StatusBadRequest
		return resp, err
	}

	result, err := compareFaces(req)
	if err != nil {
		resp.Body = err.Error()
		resp.StatusCode = http.StatusBadRequest
		return resp, err
	}

	data, err := json.Marshal(result)
	if err != nil {
		resp.Body = "Failed to encode data as JSON"
		resp.StatusCode = http.StatusBadRequest
		return resp, err
	}

	resp.Body = string(data)
	resp.StatusCode = http.StatusOK

	return resp, nil
}

func main() {
	lambda.Start(handleLambdaEvent)
}
