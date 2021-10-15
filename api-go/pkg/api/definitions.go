package api

// Response is the structure of the request response.
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
