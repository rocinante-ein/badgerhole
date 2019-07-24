package controllers

import (
	"net/http"
	"path"

	"github.com/rocinante-ein/badgerhole/internal/interfaces/presenters"
	"github.com/rocinante-ein/badgerhole/internal/usecases"
	"github.com/spf13/viper"
)

// LoggingControllerImpl all routing use this controller
type LoggingControllerImpl struct{}

// NewLoggingController Create new allcontroller instance
func NewLoggingController() *LoggingControllerImpl {
	return &LoggingControllerImpl{}
}

// HandleFunc main controller job
func (lci LoggingControllerImpl) HandleFunc(w http.ResponseWriter, r *http.Request) {

	// create logging usecase instance
	var loggingUsecase = usecases.NewLoggingUsecase()

	// logging and get wwwlog instance
	loggingUsecase.Logging(r)

	// create fake login presenter instance
	var fakeLoginPresenter = presenters.NewFakeLoginPresenter(
		path.Join(viper.GetString("TemplateDir"), "login.html"))

	// fake login view response to tanuki client
	fakeLoginPresenter.Response(w, map[string]string{
		"http_method": r.Method,
		"server":      viper.GetString("ServerName"),
	})
}
