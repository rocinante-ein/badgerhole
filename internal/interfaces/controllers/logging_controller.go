package controllers

import "net/http"

// LoggingController badger hole controller interface
type LoggingController interface {
	HandleFunc(http.ResponseWriter, *http.Request)
}
