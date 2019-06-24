package model

type HttpRequestResponse struct {
	StatusCode     int
	ResponseString string
	Body           interface{}
}
