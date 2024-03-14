package http

import "time"

type Response struct {
	Status      string
	StatusCode  int
	Message     string
	RequestTime time.Time
}
