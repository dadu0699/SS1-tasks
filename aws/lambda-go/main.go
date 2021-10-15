package main

import (
	"encoding/base64"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type compareFacesStruct struct {
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

func compareFaces(body compareFacesStruct) (*rekognition.CompareFacesOutput, error) {
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

func handleLambdaEvent(event compareFacesStruct) (response, error) {
	result, err := compareFaces(event)
	resp := response{
		Code: 200,
		Data: result,
	}

	if err != nil {
		resp = response{
			Code: 404,
			Data: err,
		}
	}

	return resp, nil
}

func main() {
	lambda.Start(handleLambdaEvent)
}
