package api

import (
	"encoding/base64"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func createSession() (*rekognition.Rekognition, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(os.Getenv("REGION_NAME")),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AWS_ACCESS_KEY_ID"),
				os.Getenv("AWS_SECRET_ACCESS_KEY"),
				""),
		},
	)
	if err != nil {
		return nil, err
	}

	svc := rekognition.New(sess)
	return svc, nil
}

// DetectLabels is a function that allows the detection of image labels through the rekognition service.
func DetectLabels(body DetectLabelsStruct) (*rekognition.DetectLabelsOutput, error) {
	decodedImage, err := base64.StdEncoding.DecodeString(body.Image)
	if err != nil {
		return nil, err
	}

	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{Bytes: decodedImage},
	}

	svc, err := createSession()
	if err != nil {
		return nil, err
	}

	result, err := svc.DetectLabels(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CompareFaces is a function that allows comparison of faces through the rekognition service.
func CompareFaces(body CompareFacesStruct) (*rekognition.CompareFacesOutput, error) {
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
