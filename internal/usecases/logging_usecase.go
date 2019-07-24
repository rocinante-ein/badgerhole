package usecases

import (
	"net/http"
)

// LoggingUsecase badger hole logging usecase
type LoggingUsecase interface {
	Logging(*http.Request)
}
