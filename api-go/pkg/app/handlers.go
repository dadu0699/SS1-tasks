package app

import (
	"api-go/pkg/api"
	"log"
	"net/http"
)

// IndexHandler handler works like healthcheck.
func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendResponse(w, r, "API running smoothly", http.StatusOK)
	}
}

// DetectLabelsHandler handler works to get the labels received from rekognition.
func (a *App) DetectLabelsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := api.DetectLabelsStruct{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(w, r, "Cannot parse post body", http.StatusBadRequest)
			return
		}

		result,err := api.DetectLabels(req)
		if err != nil {
			log.Printf("%v \n", err)
			sendResponse(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		sendResponse(w, r, result, http.StatusOK)
	}
}

// CompareFacesHandler handler works to get the array of face matches from rekognition.
func (a *App) CompareFacesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := api.CompareFacesStruct{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(w, r, "Cannot parse post body", http.StatusBadRequest)
			return
		}

		result,err := api.CompareFaces(req)
		if err != nil {
			log.Printf("%v \n", err)
			sendResponse(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		sendResponse(w, r, result, http.StatusOK)
	}
}